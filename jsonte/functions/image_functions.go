package functions

import (
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
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
		file, err := safeio.Resolver(str)
		if err != nil {
			return nil, err
		}

		config, _, err := image.DecodeConfig(file)
		if err != nil {
			return nil, err
		}
		utils.PutCache(imageCache, str, config)
		err = file.Close()
		if err != nil {
			return nil, err
		}
		img = &config
	}
	return img, nil
}

func imageWidth(str string) (utils.JsonNumber, error) {
	bounds, err := imageBounds(str)
	if err != nil {
		return utils.ToNumber(0), err
	}
	return utils.ToNumber(bounds.Width), nil
}

func imageHeight(str string) (utils.JsonNumber, error) {
	bounds, err := imageBounds(str)
	if err != nil {
		return utils.ToNumber(0), err
	}
	return utils.ToNumber(bounds.Height), nil
}
