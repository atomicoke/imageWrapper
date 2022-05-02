package main

import (
	"fmt"
	"github.com/atomicoke/imageWrapper/image"
	"net/http"
	"strconv"
)

var cache = map[string]*image.Wrap{}

func main() {
	cache = map[string]*image.Wrap{}

	addr := ":8888"
	http.Handle("/", resizer())

	fmt.Println("Server started on port http://localhost" + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func resizer() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		resizeStr := r.URL.Path[1:]
		resize, err := strconv.Atoi(resizeStr)
		if err != nil {
			http.Error(w, "Invalid resizer value", http.StatusBadRequest)
			return
		}

		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "Missing url parameter", http.StatusBadRequest)
			return
		}

		// hint cache
		key := image.BuildKey(resizeStr, url)
		if wrap, ok := cache[key]; ok {
			fmt.Println("命中缓存 : " + key)
			_, _ = wrap.WriteTo(w)
			return
		}

		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Failed to get image\ncause:"+err.Error(), http.StatusInternalServerError)
			return
		}

		wrap, err := image.NewWrap(resp.Body, resize)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("初始化   : " + key)

		wrap.FillHeader(resp.Header, "Cache-Control", "Last-Modified", "Expires", "Etag", "Link")

		cache[key] = wrap
		_, _ = wrap.WriteTo(w)
	})
}