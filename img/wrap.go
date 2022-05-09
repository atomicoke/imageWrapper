package img

import (
	"bytes"
	"errors"
	io2 "github.com/atomicoke/imageWrapper/io"
	"github.com/disintegration/imaging"
	"io"
	"net/http"
)

func NewWrap(body io.ReadCloser, m Mapping) (*Wrap, error) {
	defer body.Close()

	bs, err := io2.ReaderToBytes(body)
	if err != nil {
		return nil, errors.New("Failed to read image\ncause:" + err.Error())
	}

	// decode bytes to image
	img, err := Decode(bs)
	if err != nil {
		return nil, errors.New("Failed to decode image\ncause:" + err.Error())
	}

	// get image contentType and extension
	contentType := http.DetectContentType(bs)
	format, err := parseType(contentType)
	if err != nil {
		return nil, errors.New("Failed to parse image type\ncause:" + err.Error())
	}

	// save the resized image
	data := bytes.NewBuffer([]byte{})

	img = m(img)

	err = imaging.Encode(data, img, format)
	if err != nil {
		return nil, errors.New("Failed to encode image\ncause:" + err.Error())
	}

	return createImage(data, format, contentType).
		PutHeader("Content-Type", contentType), nil
}

func (w *Wrap) PutHeader(key, val string) *Wrap {
	w.header.Add(key, val)
	return w
}

func (w *Wrap) FillHeader(src http.Header, headerNames ...string) *Wrap {
	for _, name := range headerNames {
		k := http.CanonicalHeaderKey(name)
		for _, v := range src[k] {
			w.header.Add(k, v)
		}
	}
	return w
}

func (w Wrap) WriteTo(resp http.ResponseWriter) (int, error) {
	for s := range w.header {
		resp.Header().Set(s, w.header.Get(s))
	}

	return resp.Write(w.data.Bytes())
}
