package img

import (
	"bytes"
	"github.com/disintegration/imaging"
	"net/http"
)

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
