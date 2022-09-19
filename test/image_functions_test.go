package test

import (
	"encoding/base64"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"testing"
)

var img, _ = base64.StdEncoding.DecodeString("iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==")

func TestImageWidth(t *testing.T) {
	safeio.Resolver = safeio.CreateFakeFS(
		map[string]interface{}{
			"test.png": img,
		},
		false,
	)
	eval := evaluate(t, `imageWidth('test.png')`)
	assertNumber(t, eval, 5)
	safeio.Resolver = safeio.DefaultIOResolver
}

func TestImageHeight(t *testing.T) {
	safeio.Resolver = safeio.CreateFakeFS(
		map[string]interface{}{
			"test.png": img,
		},
		false,
	)
	eval := evaluate(t, `imageHeight('test.png')`)
	assertNumber(t, eval, 5)
	safeio.Resolver = safeio.DefaultIOResolver
}
