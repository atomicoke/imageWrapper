package image

import (
	"net/http"
)

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