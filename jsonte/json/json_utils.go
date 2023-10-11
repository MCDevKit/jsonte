package json

import (
	"bytes"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"io"
	"unicode/utf8"
)

// CacheDir is a directory used for cache
var CacheDir string

func ConvertToUTF8(str []byte) ([]byte, error) {
	if !utf8.Valid(str) {
		utils.Logger.Warnf("Input is not a valid UTF-8 string, attempting to detect and convert to UTF-8")
		r, err := chardet.NewTextDetector().DetectBest(str)
		if err != nil {
			return nil, burrito.WrapErrorf(err, "Failed to detect encoding")
		}
		byteReader := bytes.NewReader(str)
		reader, err := charset.NewReaderLabel(r.Charset, byteReader)
		if err != nil {
			return nil, burrito.WrapErrorf(err, "Failed to convert the file to UTF-8")
		}
		str, err = io.ReadAll(reader)
		if err != nil {
			return nil, burrito.WrapErrorf(err, "Failed to convert the file to UTF-8")
		}
	}
	return str, nil
}
