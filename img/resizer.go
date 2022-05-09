package img

import (
	"github.com/disintegration/imaging"
	"image"
	"net/http"
	"strconv"
	"strings"
)

func Resizer() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var width, height, resize int
		var err error

		// prepare start
		resizeStr := r.URL.Path[1:]
		split := strings.Split(resizeStr, "x")
		if len(split) < 2 {
			resize, err = strconv.Atoi(resizeStr)
			if err != nil {
				http.Error(w, "Invalid resize value", http.StatusBadRequest)
				return
			}

			width = resize
			height = 0
		} else {
			widthStr := split[0]
			heightStr := split[1]
			width, err = strconv.Atoi(widthStr)
			if err != nil {
				http.Error(w, "Invalid width value", http.StatusBadRequest)
				return
			}

			height, err = strconv.Atoi(heightStr)
			if err != nil {
				http.Error(w, "Invalid height value", http.StatusBadRequest)
				return
			}
		}

		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "Missing url parameter", http.StatusBadRequest)
			return
		}
		// prepare end

		// hint cache
		key := BuildKey(resizeStr, url)
		if wrap, ok := GetFromCache(key); ok {
			_, _ = wrap.WriteTo(w)
			return
		}

		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Failed to get image\ncause:"+err.Error(), http.StatusInternalServerError)
			return
		}

		// resize image
		wrap, err := NewWrap(resp.Body, func(src image.Image) image.Image {
			return imaging.Resize(src, width, height, imaging.Lanczos)
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		wrap.FillHeader(resp.Header, "Cache-Control", "Last-Modified", "Expires", "Etag", "Link")
		wrap.PutHeaderNx("Cache-Control", "max-age=888888888")

		PutToCache(key, wrap)
		_, _ = wrap.WriteTo(w)
	})
}
