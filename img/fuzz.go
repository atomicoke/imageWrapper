package img

import (
	"github.com/disintegration/imaging"
	"image"
	"net/http"
)

const (
	blur    = 3
	process = "blur"
)

func Fuzz() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		url := req.URL.Query().Get("url")
		if url == "" {
			return
		}

		key := BuildKey(process, url)
		if wrap, ok := GetFromCache(key); ok {
			_, _ = wrap.WriteTo(w)
			return
		}

		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Failed to get image\ncause:"+err.Error(), http.StatusInternalServerError)
			return
		}

		wrap, err := NewWrap(resp.Body, func(src image.Image) image.Image {
			return imaging.Blur(src, blur)
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		wrap.FillHeader(resp.Header, "Cache-Control", "Last-Modified", "Expires", "Etag", "Link")
		wrap.PutHeaderNx("Cache-Control", "public, max-age=3153600000")

		PutToCache(key, wrap)
		_, _ = wrap.WriteTo(w)
	}
}
