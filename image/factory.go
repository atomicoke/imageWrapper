package image

import (
	"bytes"
	"errors"
	"github.com/disintegration/imaging"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"net/http"
)

func NewWrap(body io.ReadCloser, width, height int) (*Wrap, error) {
	defer body.Close()

	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, errors.New("Failed to read image\ncause:" + err.Error())
	}

	//img, name, err := image.Decode(bytes.NewReader(bs))
	img, err := Decode(bs)
	if err != nil {
		return nil, errors.New("Failed to decode image\ncause:" + err.Error())
	}

	img = imaging.Resize(img, width, height, imaging.Lanczos)
	data := bytes.NewBuffer([]byte{})

	contentType := http.DetectContentType(bs)

	format, err := parseType(contentType)
	if err != nil {
		return nil, errors.New("Failed to parse image type\ncause:" + err.Error())
	}
	err = imaging.Encode(data, img, format)
	if err != nil {
		return nil, errors.New("Failed to encode image\ncause:" + err.Error())
	}
	return createImage(data, format, contentType).
		PutHeader("Content-Type", contentType), nil
}

func BuildKey(resizeStr string, imageUrl string) string {
	return resizeStr + " - " + imageUrl
}

func createImage(data *bytes.Buffer, format imaging.Format, contentType string) *Wrap {
	return &Wrap{
		data:        data,
		header:      http.Header{},
		format:      format,
		contentType: contentType,
	}
}

func parseType(contentType string) (imaging.Format, error) {
	switch contentType {
	case "image/jpeg":
		return imaging.JPEG, nil
	case "image/png":
		return imaging.PNG, nil
	case "image/gif":
		return imaging.GIF, nil
	case "image/bmp":
		return imaging.BMP, nil
	case "image/tiff":
		return imaging.TIFF, nil
	case "image/webp":
		return imaging.JPEG, nil
	default:
		return imaging.JPEG, nil
	}
}