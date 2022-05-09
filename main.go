package main

import (
	"flag"
	"github.com/atomicoke/imageWrapper/img"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

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

	http.Handle("/", img.Resizer())
	http.Handle("/fuzz", img.Fuzz())

	log.Infoln("Server started on port http://localhost" + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
