package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var cache = map[string][]byte{}

func main() {
	cache = map[string][]byte{}

	addr := ":8888"
	http.Handle("/", handler())

	fmt.Println("Server started on port http://localhost" + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}

}

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		resizeStr := r.URL.Path[1:]
		_, err := strconv.Atoi(resizeStr)
		if err != nil {
			http.Error(w, "Invalid resize value", http.StatusBadRequest)
			return
		}

		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "Missing url parameter", http.StatusBadRequest)
			return
		}

		key := url + ":" + resizeStr
		data := cache[key]
		if !(len(data) == 0 || data == nil) {
			w.Write(data)
			return
		}

		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Failed to get image\ncause:"+err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		copyHeader(w.Header(), resp.Header, "Cache-Control", "Last-Modified", "Expires", "Etag", "Link")

		data, _ = ioutil.ReadAll(resp.Body)

		cache[key] = data
		// todo 记录请求头到map
		// todo resize
		w.Write(data)
	})
}

func copyHeader(dst, src http.Header, headerNames ...string) {
	for _, name := range headerNames {
		k := http.CanonicalHeaderKey(name)
		for _, v := range src[k] {
			dst.Add(k, v)
		}
	}
}