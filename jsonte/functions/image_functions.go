package functions

import (
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

const imageCache = "imageBounds"

func RegisterImageFunctions() {
	utils.CreateCacheBucket(imageCache, 3600)
	RegisterFunction(JsonFunction{
		Name: "imageWidth",
		Body: imageWidth,
	})
	RegisterFunction(JsonFunction{
		Name: "imageHeight",
		Body: imageHeight,
	})
}

func imageBounds(str string) (*image.Config, error) {
	var img *image.Config
	cache := utils.GetCache(imageCache, str)
	if cache != nil {
		c := (*cache).(image.Config)
		img = &c
	} else {
		file, err := safeio.Resolver.Open(str)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, utils.WrappedErrorf("File '%s' does not exist", str)
			}
			return nil, utils.WrapErrorf(err, "Failed to open image file %s", str)
		}

		config, _, err := image.DecodeConfig(file)
		if err != nil {
			return nil, utils.WrapErrorf(err, "Failed to decode image file %s", str)
		}
		utils.PutCache(imageCache, str, config)
		err = file.Close()
		if err != nil {
			return nil, utils.WrapErrorf(err, "Failed to close image file %s", str)
		}
		img = &config
	}
	return img, nil
}

func imageWidth(str string) (utils.JsonNumber, error) {
	bounds, err := imageBounds(str)
	if err != nil {
		return utils.ToNumber(0), utils.WrapErrorf(err, "Failed to get image bounds for %s", str)
	}
	return utils.ToNumber(bounds.Width), nil
}

func imageHeight(str string) (utils.JsonNumber, error) {
	bounds, err := imageBounds(str)
	if err != nil {
		return utils.ToNumber(0), utils.WrapErrorf(err, "Failed to get image bounds for %s", str)
	}
	return utils.ToNumber(bounds.Height), nil
}
