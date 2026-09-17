// Package brarchive reads Minecraft Bedrock .brarchive files.
//
// Newer Bedrock versions ship vanilla pack directories as archives stored under
// <pack>/__brarchive/<dir path>.brarchive. Each archive holds only the files
// directly inside that directory (no subdirectories). The format is:
//
//	8 bytes		magic
//	uint32 LE	number of entries
//	uint32 LE	version (1)
//	entries		256 bytes each: uint8 name length, name (zero padded to 247 bytes),
//				uint32 LE offset, uint32 LE size
//	data		concatenated file contents, offsets are relative to the data start
//
// Empty files have offset 0 and size 0.
//
// A file inside an archive is addressed with a pseudo path of the form
// <archive>.brarchive#/<name>, see SplitPath.
package brarchive

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/Bedrock-OSS/go-burrito/burrito"
	"golang.org/x/exp/mmap"
)

const Extension = ".brarchive"

// Separator separates the archive path from the entry name in a pseudo path.
const Separator = "#/"

const headerSize = 16
const entrySize = 256

var magic = []byte{0x7d, 0x27, 0x25, 0xb1, 0xa0, 0x52, 0x70, 0x26}

type entry struct {
	offset int64
	size   int64
}

// Archive is a parsed .brarchive file.
type Archive struct {
	reader  io.ReaderAt
	size    int64
	modTime time.Time
	entries map[string]entry
	names   []string
}

var (
	mu    sync.RWMutex
	cache = map[string]*Archive{}
)

// SplitPath splits a pseudo path <archive>.brarchive#/<name> into the archive
// path and the entry name. ok is false if p is not such a pseudo path.
func SplitPath(p string) (archive, name string, ok bool) {
	i := strings.Index(p, Extension+"#")
	if i < 0 {
		return "", "", false
	}
	archive = p[:i+len(Extension)]
	name = strings.TrimLeft(p[i+len(Extension)+1:], "/\\")
	if name == "" {
		return "", "", false
	}
	return archive, name, true
}

// JoinPath builds a pseudo path for the entry name inside the archive.
func JoinPath(archive, name string) string {
	return archive + Separator + name
}

// Open memory-maps and parses the archive. Archives are cached for the lifetime
// of the process, so the returned Archive must not be closed by the caller.
func Open(p string) (*Archive, error) {
	key := filepath.Clean(p)
	mu.RLock()
	a, ok := cache[key]
	mu.RUnlock()
	if ok {
		return a, nil
	}
	mu.Lock()
	defer mu.Unlock()
	if a, ok := cache[key]; ok {
		return a, nil
	}
	a, err := open(p)
	if err != nil {
		return nil, err
	}
	cache[key] = a
	return a, nil
}

func open(p string) (*Archive, error) {
	r, err := mmap.Open(p)
	if err != nil {
		return nil, err
	}
	// mmap.Open already stats the file; the modification time is not worth a second syscall
	a, err := Parse(r, int64(r.Len()), time.Time{})
	if err != nil {
		_ = r.Close()
		return nil, burrito.WrapErrorf(err, "Failed to parse archive %s", p)
	}
	return a, nil
}

// Parse parses an archive of the given size from r. The reader is kept for reading entries.
func Parse(r io.ReaderAt, size int64, modTime time.Time) (*Archive, error) {
	a := &Archive{reader: r, size: size, modTime: modTime, entries: map[string]entry{}}
	if err := a.parse(); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Archive) parse() error {
	header := make([]byte, headerSize)
	if _, err := a.reader.ReadAt(header, 0); err != nil {
		return burrito.WrapErrorf(err, "Archive is too short")
	}
	if !bytes.Equal(header[:8], magic) {
		return burrito.WrappedErrorf("Invalid archive magic")
	}
	count := int64(binary.LittleEndian.Uint32(header[8:]))
	version := binary.LittleEndian.Uint32(header[12:])
	if version != 1 {
		return burrito.WrappedErrorf("Unsupported archive version %d", version)
	}
	dataStart := headerSize + count*entrySize
	if dataStart > a.size {
		return burrito.WrappedErrorf("Archive is too short for %d entries", count)
	}
	a.entries = make(map[string]entry, count)
	a.names = make([]string, 0, count)
	buf := make([]byte, entrySize)
	for i := int64(0); i < count; i++ {
		if _, err := a.reader.ReadAt(buf, headerSize+i*entrySize); err != nil {
			return burrito.WrapErrorf(err, "Failed to read entry %d", i)
		}
		nameLen := int(buf[0])
		if nameLen > entrySize-9 {
			return burrito.WrappedErrorf("Invalid name length %d in entry %d", nameLen, i)
		}
		name := string(buf[1 : 1+nameLen])
		e := entry{
			offset: dataStart + int64(binary.LittleEndian.Uint32(buf[entrySize-8:])),
			size:   int64(binary.LittleEndian.Uint32(buf[entrySize-4:])),
		}
		if e.offset+e.size > a.size {
			return burrito.WrappedErrorf("Entry %s points outside of the archive", name)
		}
		a.entries[name] = e
		a.names = append(a.names, name)
	}
	return nil
}

// Names returns the names of all files in the archive.
func (a *Archive) Names() []string {
	return a.names
}

// Open returns a reader for the named file.
func (a *Archive) Open(name string) (*io.SectionReader, error) {
	e, ok := a.entries[name]
	if !ok {
		return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrNotExist}
	}
	return io.NewSectionReader(a.reader, e.offset, e.size), nil
}

// Stat returns file info for the named file.
func (a *Archive) Stat(name string) (fs.FileInfo, error) {
	e, ok := a.entries[name]
	if !ok {
		return nil, &fs.PathError{Op: "stat", Path: name, Err: fs.ErrNotExist}
	}
	return &fileInfo{name: name, size: e.size, modTime: a.modTime}, nil
}

type fileInfo struct {
	name    string
	size    int64
	modTime time.Time
}

func (f *fileInfo) Name() string       { return f.name }
func (f *fileInfo) Size() int64        { return f.size }
func (f *fileInfo) Mode() fs.FileMode  { return 0444 }
func (f *fileInfo) ModTime() time.Time { return f.modTime }
func (f *fileInfo) IsDir() bool        { return false }
func (f *fileInfo) Sys() interface{}   { return nil }
