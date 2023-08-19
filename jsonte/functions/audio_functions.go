package functions

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
	"io"
	"path/filepath"
	"strings"
)

const audioCache = "audioLength"

func RegisterAudioFunctions() {
	const group = "audio"
	RegisterGroup(Group{
		Name:    group,
		Title:   "Audio functions",
		Summary: "Audio functions are related to reading various information about audio files.",
	})
	utils.CreateCacheBucket(audioCache)
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "audioDuration",
		Body:  audioDuration,
		Docs: Docs{
			Summary: "Returns the duration of an audio file in seconds.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the audio file.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "test": "{{audioDuration('resources/sounds/sound.wav')}}"
  }
}
</code>`,
		},
	})
}

func audioDuration(str *types.JsonString) (*types.JsonNumber, error) {
	var length float64 = 0
	cache := utils.GetCache(audioCache, str.StringValue())
	if cache != nil {
		length = (*cache).(float64)
	} else {
		file, err := safeio.Resolver.Open(str.StringValue())
		if err != nil {
			return types.AsNumber(0), burrito.WrapErrorf(err, "Failed to open audio file %s", str.StringValue())
		}
		streamer, format, err := decodeAudio(str.StringValue(), file)
		if err != nil {
			return types.AsNumber(0), burrito.WrapErrorf(err, "Failed to decode audio file %s", str.StringValue())
		}
		streamer.Len()
		length = float64(streamer.Len()) / float64(format.SampleRate)
		utils.PutCache(audioCache, str.StringValue(), length)
		err = file.Close()
		if err != nil {
			return types.AsNumber(0), burrito.WrapErrorf(err, "Failed to close audio file %s", str.StringValue())
		}
	}
	return types.AsNumber(length), nil
}

func decodeAudio(path string, file io.ReadCloser) (beep.StreamSeekCloser, beep.Format, error) {
	if strings.HasSuffix(path, ".ogg") {
		return vorbis.Decode(file)
	}
	if strings.HasSuffix(path, ".mp3") {
		return mp3.Decode(file)
	}
	if strings.HasSuffix(path, ".wav") {
		return wav.Decode(file)
	}
	return nil, beep.Format{}, burrito.WrappedErrorf("Unsupported audio file format %s", filepath.Ext(path))
}
