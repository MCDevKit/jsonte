package functions

import (
	"errors"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
	"jsonte/jsonte/io"
	"jsonte/jsonte/utils"
	"os"
	"strings"
)

const audioCache = "audioLength"

func RegisterAudioFunctions() {
	utils.CreateCacheBucket(audioCache, 3600)
	RegisterFunction(JsonFunction{
		Name: "audioDuration",
		Body: audioDuration,
	})
}

func audioDuration(str string) (utils.JsonNumber, error) {
	// audioDuration('C:\Users\brzoz\Downloads\COMCell_Message 1 (ID 1111)_BSB.ogg')
	// audioDuration('C:\MCLauncher\imported_versions\Minecraft.Windows_1.19.2022.0_x64.appx\data\resource_packs\vanilla\sounds\dig\ancient_debris2.fsb')
	var length float64 = 0
	cache := utils.GetCache(audioCache, str)
	if cache != nil {
		length = (*cache).(float64)
	} else {
		file, err := io.Resolver.Resolve(str)
		if err != nil {
			return utils.ToNumber(0), err
		}
		streamer, format, err := decodeAudio(str, file)
		if err != nil {
			return utils.ToNumber(0), err
		}
		streamer.Len()
		length = float64(streamer.Len()) / float64(format.SampleRate)
		utils.PutCache(audioCache, str, length)
		err = file.Close()
		if err != nil {
			return utils.ToNumber(0), err
		}
	}
	return utils.ToNumber(length), nil
}

func decodeAudio(path string, file *os.File) (beep.StreamSeekCloser, beep.Format, error) {
	if strings.HasSuffix(path, ".ogg") {
		return vorbis.Decode(file)
	}
	if strings.HasSuffix(path, ".mp3") {
		return mp3.Decode(file)
	}
	if strings.HasSuffix(path, ".wav") {
		return wav.Decode(file)
	}
	return nil, beep.Format{}, errors.New("unsupported audio format")
}
