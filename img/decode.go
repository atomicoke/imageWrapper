package img

import (
	"bytes"
	"errors"
	_ "golang.org/x/image/bmp"
	"golang.org/x/image/webp"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func Decode(data []byte) (image.Image, error) {
	if data == nil || len(data) == 0 {
		return nil, errors.New("data is empty")
	}

	img, imageType, err := image.Decode(bytes.NewReader(data))

	switch imageType {
	case `jpeg`:
		return img, nil
	case `png`:
		return img, nil
	case `gif`:
		return img, nil
	case `bmp`:
		return img, nil
	default:
		if err == nil {
			return img, nil
		}

		// 尝试以 webp 进行解码
		img, err = webp.Decode(bytes.NewReader(data))
		if err == nil {
			return img, nil
		}
	}
	return nil, err
}
