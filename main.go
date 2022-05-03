package main

import (
	"flag"
	"github.com/atomicoke/imageWrapper/image"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

var cache = map[string]*image.Wrap{}

var port = flag.Int("p", 8888, "port to listen on")
var debug = flag.Bool("d", false, "log level use debug")
var usage = flag.Bool("u", false, "show usage")

func init() {
	flag.Parse()

	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "time",
			log.FieldKeyLevel: "level",
			log.FieldKeyMsg:   "message",
		},
		TimestampFormat: "2006-01-02 15:04:05 111",
		PrettyPrint:     true,
	})

	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

}

func main() {
	if *usage {
		flag.Usage()
		return
	}

	addr := ":" + strconv.Itoa(*port)

	cache = map[string]*image.Wrap{}

	http.Handle("/", resizer())

	log.Infoln("Server started on port http://localhost" + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func resizer() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var width, height, resize int
		var err error

		resizeStr := r.URL.Path[1:]
		split := strings.Split(resizeStr, "x")
		if len(split) < 2 {
			resize, err = strconv.Atoi(resizeStr)
			if err != nil {
				http.Error(w, "Invalid resize value", http.StatusBadRequest)
				return
			}

			width = resize
			height = resize
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

		// hint cache
		key := image.BuildKey(resizeStr, url)
		if wrap, ok := cache[key]; ok {
			log.Debugln("命中缓存 : " + key)
			_, _ = wrap.WriteTo(w)
			return
		}

		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Failed to get image\ncause:"+err.Error(), http.StatusInternalServerError)
			return
		}

		wrap, err := image.NewWrap(resp.Body, width, height)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Debugln("初始化   : " + key)

		wrap.FillHeader(resp.Header, "Cache-Control", "Last-Modified", "Expires", "Etag", "Link")

		cache[key] = wrap
		_, _ = wrap.WriteTo(w)
	})
}