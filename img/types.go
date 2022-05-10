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
	// the content type of image
	// e.g. image/png
	contentType string
}

// Mapping is a conversion operation,from src to dest.
// e.g resize,crop,rotate,flip,flop,blur,sharpen,grayscale,invert,convolve,blend,composite
type Mapping func(src image.Image) (dest image.Image)

// Cache use to save the image
type Cache interface {
	Put(key string, wrap *Wrap)

	Get(key string) (wrap *Wrap, ok bool)
}

type CacheModel string

const (
	// MEMORY  内存实现
	MEMORY CacheModel = "memory"
)

type MemoryCache struct {
	cache map[string]*Wrap
}
