package handler

import (
	"net/http"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Metods", "GET,POST,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Request-With, Accept")
		next.ServeHTTP(w, r)
	})
}
