package image

import (
	"bytes"
	"fmt"
	_ "golang.org/x/image/bmp"
	"golang.org/x/image/webp"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func Decode(data []byte) (image.Image, error) {
	img, imageType, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	switch imageType {
	case `jpeg`:
	case `png`:
	case `gif`:
	case `bmp`:
		return img, nil
	default:
		// 尝试以 webp 进行解码
		img, err = webp.Decode(bytes.NewReader(data))
		fmt.Println("webp decode")
		if err == nil {
			return img, nil
		}
	}
	return nil, err
}