package img

import "net/http"

func Fuzz() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
