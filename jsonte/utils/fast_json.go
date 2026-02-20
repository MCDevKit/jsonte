package utils

import (
	"bytes"
	"sort"
	"strconv"

	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/zeebo/xxh3"
)

// Kind describes the JSON node type.
type Kind uint8

const (
	Object Kind = iota
	Array
	String
	Number
	Boolean
	Null
)

// Node points into the original source bytes.
// KeyLo/KeyHi are set only for object members; for array elements they are -1.
type Node struct {
	Kind   Kind
	Parent int32

	// Direct children live in Index.ChildRefs[ChildLo:ChildHi].
	ChildLo int32
	ChildHi int32

	// ValLo/ValHi points to the exact value bytes in source.
	ValLo int32
	ValHi int32

	// Key span without surrounding quotes.
	KeyLo int32
	KeyHi int32
}

// Index is a structural representation of a JSON document without materializing values.
type Index struct {
	Src       []byte
	Nodes     []Node
	ChildRefs []int32

	keyTables map[int][]KeyTableEntry
}

// Build parses JSON/JSONC bytes into an Index.
func Build(src []byte) (*Index, error) {
	p := &parser{
		src: src,
		idx: &Index{Src: src},
	}
	if err := p.parse(); err != nil {
		return nil, burrito.WrapError(err, "Failed to parse JSON")
	}
	return p.idx, nil
}

// Children returns direct child node indices for node n.
// It returns nil for leaves and out-of-range indices.
func (idx *Index) Children(n int) []int {
	if n < 0 || n >= len(idx.Nodes) {
		return nil
	}
	node := idx.Nodes[n]
	if node.ChildLo == node.ChildHi {
		return nil
	}
	children := make([]int, node.ChildHi-node.ChildLo)
	for i := range children {
		children[i] = int(idx.ChildRefs[int(node.ChildLo)+i])
	}
	return children
}

// KeyTableEntry stores preprocessed object member metadata.
// Hash is computed from the unescaped key bytes.
type KeyTableEntry struct {
	Hash uint64

	// Raw key span (without quotes, escapes still present in source bytes).
	KeyLo int32
	KeyHi int32

	ValIndex  int32
	HasDollar bool
	HasCaret  bool
	HasEscape bool
}

// BuildKeyTable returns object members sorted by (hash, canonical-key-bytes).
// If buf is nil, it may return an internal cached slice.
func (idx *Index) BuildKeyTable(objIdx int, buf []KeyTableEntry) ([]KeyTableEntry, error) {
	if objIdx < 0 || objIdx >= len(idx.Nodes) || idx.Nodes[objIdx].Kind != Object {
		return nil, burrito.WrappedError("BuildKeyTable: Not an object")
	}
	if cached, ok := idx.keyTables[objIdx]; ok {
		if buf == nil {
			return cached, nil
		}
		buf = buf[:0]
		buf = append(buf, cached...)
		return buf, nil
	}

	type buildEntry struct {
		entry KeyTableEntry
		key   []byte
	}

	obj := idx.Nodes[objIdx]
	entries := make([]buildEntry, 0, obj.ChildHi-obj.ChildLo)
	for childRef := obj.ChildLo; childRef < obj.ChildHi; childRef++ {
		valIdx := idx.ChildRefs[childRef]
		child := idx.Nodes[valIdx]
		if child.KeyLo < 0 || child.KeyHi < 0 {
			continue
		}

		rawKey := idx.Src[child.KeyLo:child.KeyHi]
		canonKey := rawKey
		hasEscape := bytes.IndexByte(rawKey, '\\') >= 0
		if hasEscape {
			decoded, err := unescapeJSONStringContent(rawKey)
			if err != nil {
				return nil, burrito.WrapError(err, "BuildKeyTable: Failed to unescape object key")
			}
			canonKey = decoded
		}

		entries = append(entries, buildEntry{
			entry: KeyTableEntry{
				Hash:      xxh3.Hash(canonKey),
				KeyLo:     child.KeyLo,
				KeyHi:     child.KeyHi,
				ValIndex:  valIdx,
				HasDollar: len(canonKey) > 0 && canonKey[0] == '$',
				HasCaret:  len(canonKey) > 0 && canonKey[0] == '^',
				HasEscape: hasEscape,
			},
			key: canonKey,
		})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].entry.Hash != entries[j].entry.Hash {
			return entries[i].entry.Hash < entries[j].entry.Hash
		}
		return bytes.Compare(entries[i].key, entries[j].key) < 0
	})

	table := make([]KeyTableEntry, len(entries))
	for i := range entries {
		table[i] = entries[i].entry
	}

	if idx.keyTables == nil {
		idx.keyTables = make(map[int][]KeyTableEntry)
	}
	idx.keyTables[objIdx] = table

	if buf == nil {
		return table, nil
	}
	buf = buf[:0]
	buf = append(buf, table...)
	return buf, nil
}

type parser struct {
	src []byte
	pos int
	idx *Index

	childrenByNode [][]int32
}

func (p *parser) parse() error {
	p.skipWS()
	if p.pos >= len(p.src) {
		return burrito.WrappedError("Empty input")
	}

	if _, err := p.readValue(-1, -1, -1); err != nil {
		return burrito.WrapError(err, "Failed to read root value")
	}

	p.skipWS()
	if p.pos != len(p.src) {
		return burrito.WrappedError("Trailing bytes found after root value")
	}

	p.materializeChildren()
	return nil
}

func (p *parser) readValue(parent int, keyLo, keyHi int32) (int, error) {
	p.skipWS()
	if p.pos >= len(p.src) {
		return -1, burrito.WrappedError("Unexpected EOF")
	}

	switch c := p.src[p.pos]; c {
	case '{':
		return p.readObject(parent)
	case '[':
		return p.readArray(parent)
	case '"':
		return p.readString(parent, keyLo, keyHi)
	case 't', 'f':
		return p.readBoolean(parent, keyLo, keyHi)
	case 'n':
		return p.readNull(parent, keyLo, keyHi)
	default:
		if c == '-' || (c >= '0' && c <= '9') {
			return p.readNumber(parent, keyLo, keyHi)
		}
	}

	return -1, burrito.WrappedError("Invalid value")
}

func (p *parser) readObject(parent int) (int, error) {
	start := p.pos
	p.pos++

	nodeIdx := p.newNode(Node{Kind: Object, Parent: int32(parent), ValLo: int32(start)})

	p.skipWS()
	if p.pos < len(p.src) && p.src[p.pos] == '}' {
		p.pos++
		p.idx.Nodes[nodeIdx].ValHi = int32(p.pos)
		return nodeIdx, nil
	}

	for {
		p.skipWS()
		if p.pos >= len(p.src) || p.src[p.pos] != '"' {
			return -1, burrito.WrappedError("Object key must be a string")
		}

		klo, khi, err := p.scanString()
		if err != nil {
			return -1, burrito.WrapError(err, "Failed to read object key")
		}

		p.skipWS()
		if p.pos >= len(p.src) || p.src[p.pos] != ':' {
			return -1, burrito.WrappedError("Expected ':' after object key")
		}
		p.pos++

		valIdx, err := p.readValue(nodeIdx, int32(klo), int32(khi))
		if err != nil {
			return -1, burrito.WrapError(err, "Failed to read object value")
		}

		child := &p.idx.Nodes[valIdx]
		child.KeyLo = int32(klo)
		child.KeyHi = int32(khi)
		p.addChild(nodeIdx, valIdx)

		p.skipWS()
		if p.pos >= len(p.src) {
			return -1, burrito.WrappedError("Unexpected EOF while reading object")
		}
		if p.src[p.pos] == '}' {
			p.pos++
			p.idx.Nodes[nodeIdx].ValHi = int32(p.pos)
			return nodeIdx, nil
		}
		if p.src[p.pos] != ',' {
			return -1, burrito.WrappedError("Expected ',' in object")
		}
		p.pos++
	}
}

func (p *parser) readArray(parent int) (int, error) {
	start := p.pos
	p.pos++

	nodeIdx := p.newNode(Node{Kind: Array, Parent: int32(parent), ValLo: int32(start)})

	p.skipWS()
	if p.pos < len(p.src) && p.src[p.pos] == ']' {
		p.pos++
		p.idx.Nodes[nodeIdx].ValHi = int32(p.pos)
		return nodeIdx, nil
	}

	for {
		childIdx, err := p.readValue(nodeIdx, -1, -1)
		if err != nil {
			return -1, burrito.WrapError(err, "Failed to read array element")
		}
		p.addChild(nodeIdx, childIdx)

		p.skipWS()
		if p.pos >= len(p.src) {
			return -1, burrito.WrappedError("Unexpected EOF while reading array")
		}
		if p.src[p.pos] == ']' {
			p.pos++
			p.idx.Nodes[nodeIdx].ValHi = int32(p.pos)
			return nodeIdx, nil
		}
		if p.src[p.pos] != ',' {
			return -1, burrito.WrappedError("Expected ',' in array")
		}
		p.pos++
	}
}

func (p *parser) readString(parent int, keyLo, keyHi int32) (int, error) {
	valLo, valHi, err := p.scanString()
	if err != nil {
		return -1, burrito.WrapError(err, "Failed to read string value")
	}

	return p.newNode(Node{
		Kind:   String,
		Parent: int32(parent),
		ValLo:  int32(valLo),
		ValHi:  int32(valHi),
		KeyLo:  keyLo,
		KeyHi:  keyHi,
	}), nil
}

func (p *parser) readNumber(parent int, keyLo, keyHi int32) (int, error) {
	start := p.pos

	if p.src[p.pos] == '-' {
		p.pos++
		if p.pos >= len(p.src) {
			return -1, burrito.WrappedError("Dangling '-' in number")
		}
	}

	if p.src[p.pos] == '0' {
		p.pos++
	} else {
		if p.src[p.pos] < '0' || p.src[p.pos] > '9' {
			return -1, burrito.WrappedError("Invalid number")
		}
		for p.pos < len(p.src) && p.src[p.pos] >= '0' && p.src[p.pos] <= '9' {
			p.pos++
		}
	}

	if p.pos < len(p.src) && p.src[p.pos] == '.' {
		p.pos++
		if p.pos >= len(p.src) || p.src[p.pos] < '0' || p.src[p.pos] > '9' {
			return -1, burrito.WrappedError("Invalid fraction")
		}
		for p.pos < len(p.src) && p.src[p.pos] >= '0' && p.src[p.pos] <= '9' {
			p.pos++
		}
	}

	if p.pos < len(p.src) && (p.src[p.pos] == 'e' || p.src[p.pos] == 'E') {
		p.pos++
		if p.pos < len(p.src) && (p.src[p.pos] == '+' || p.src[p.pos] == '-') {
			p.pos++
		}
		if p.pos >= len(p.src) || p.src[p.pos] < '0' || p.src[p.pos] > '9' {
			return -1, burrito.WrappedError("Invalid exponent")
		}
		for p.pos < len(p.src) && p.src[p.pos] >= '0' && p.src[p.pos] <= '9' {
			p.pos++
		}
	}

	return p.newNode(Node{
		Kind:   Number,
		Parent: int32(parent),
		ValLo:  int32(start),
		ValHi:  int32(p.pos),
		KeyLo:  keyLo,
		KeyHi:  keyHi,
	}), nil
}

func (p *parser) readBoolean(parent int, keyLo, keyHi int32) (int, error) {
	start := p.pos

	if hasTokenPrefix(p.src[p.pos:], "true") {
		p.pos += 4
		return p.newNode(Node{Kind: Boolean, Parent: int32(parent), ValLo: int32(start), ValHi: int32(p.pos), KeyLo: keyLo, KeyHi: keyHi}), nil
	}
	if hasTokenPrefix(p.src[p.pos:], "false") {
		p.pos += 5
		return p.newNode(Node{Kind: Boolean, Parent: int32(parent), ValLo: int32(start), ValHi: int32(p.pos), KeyLo: keyLo, KeyHi: keyHi}), nil
	}

	return -1, burrito.WrappedError("Invalid boolean value")
}

func (p *parser) readNull(parent int, keyLo, keyHi int32) (int, error) {
	start := p.pos
	if !hasTokenPrefix(p.src[p.pos:], "null") {
		return -1, burrito.WrappedError("Invalid null value")
	}
	p.pos += 4
	return p.newNode(Node{Kind: Null, Parent: int32(parent), ValLo: int32(start), ValHi: int32(p.pos), KeyLo: keyLo, KeyHi: keyHi}), nil
}

// scanString reads a JSON string and returns the content span (without quotes).
func (p *parser) scanString() (lo, hi int, err error) {
	if p.pos >= len(p.src) || p.src[p.pos] != '"' {
		return 0, 0, burrito.WrappedError("ScanString: Not positioned at '\"'")
	}

	p.pos++
	lo = p.pos
	for p.pos < len(p.src) {
		ch := p.src[p.pos]
		if ch == '"' {
			hi = p.pos
			p.pos++
			return lo, hi, nil
		}
		if ch == '\\' {
			p.pos++
			if p.pos >= len(p.src) {
				return 0, 0, burrito.WrappedError("Unterminated escape sequence")
			}
			if p.src[p.pos] == 'u' {
				if p.pos+4 >= len(p.src) {
					return 0, 0, burrito.WrappedError("Incomplete \\u escape sequence")
				}
				p.pos += 4
			}
		}
		p.pos++
	}

	return 0, 0, burrito.WrappedError("Unterminated string literal")
}

// skipWS skips whitespace and JSONC comments (// and /* */).
func (p *parser) skipWS() {
	for p.pos < len(p.src) {
		switch p.src[p.pos] {
		case ' ', '\t', '\r', '\n':
			p.pos++
		case '/':
			if p.pos+1 >= len(p.src) {
				return
			}
			next := p.src[p.pos+1]
			if next == '/' {
				p.pos += 2
				for p.pos < len(p.src) && p.src[p.pos] != '\n' && p.src[p.pos] != '\r' {
					p.pos++
				}
				continue
			}
			if next == '*' {
				p.pos += 2
				closed := false
				for p.pos+1 < len(p.src) {
					if p.src[p.pos] == '*' && p.src[p.pos+1] == '/' {
						p.pos += 2
						closed = true
						break
					}
					p.pos++
				}
				if closed {
					continue
				}
				p.pos = len(p.src)
				return
			}
			return
		default:
			return
		}
	}
}

func (p *parser) newNode(n Node) int {
	idx := len(p.idx.Nodes)
	p.idx.Nodes = append(p.idx.Nodes, n)
	p.childrenByNode = append(p.childrenByNode, nil)
	return idx
}

func (p *parser) addChild(parentNode, childNode int) {
	p.childrenByNode[parentNode] = append(p.childrenByNode[parentNode], int32(childNode))
}

func (p *parser) materializeChildren() {
	p.idx.ChildRefs = p.idx.ChildRefs[:0]
	for nodeIdx := range p.idx.Nodes {
		children := p.childrenByNode[nodeIdx]
		p.idx.Nodes[nodeIdx].ChildLo = int32(len(p.idx.ChildRefs))
		p.idx.ChildRefs = append(p.idx.ChildRefs, children...)
		p.idx.Nodes[nodeIdx].ChildHi = int32(len(p.idx.ChildRefs))
	}
}

func hasTokenPrefix(src []byte, token string) bool {
	return len(src) >= len(token) && string(src[:len(token)]) == token
}

func unescapeJSONStringContent(raw []byte) ([]byte, error) {
	if bytes.IndexByte(raw, '\\') < 0 {
		return raw, nil
	}

	quoted := make([]byte, 0, len(raw)+2)
	quoted = append(quoted, '"')
	quoted = append(quoted, raw...)
	quoted = append(quoted, '"')

	unquoted, err := strconv.Unquote(string(quoted))
	if err != nil {
		return nil, err
	}
	return []byte(unquoted), nil
}
