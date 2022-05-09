package img

import (
	"bytes"
	"github.com/disintegration/imaging"
	"image"
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

// Mapping is a conversion operation,from src to dest.
// e.g resize,crop,rotate,flip,flop,blur,sharpen,grayscale,invert,convolve,blend,composite
type Mapping func(src image.Image) (dest image.Image)
