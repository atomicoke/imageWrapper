package image

import (
	"bytes"
	"github.com/disintegration/imaging"
	"net/http"
)

type Wrap struct {
	// save image data
	data *bytes.Buffer
	// the header will return to client
	header http.Header
	// image type
	format imaging.Format
	// image width and height
	resize int
	// the content type of image
	// e.g. image/png
	contentType string
}