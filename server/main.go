package main

import (
	"flag"
	"fmt"
	"ipssi/handlers"
	"net/http"

	"github.com/nanoninja/bulma"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my Homepage."))

}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aboute us.")

}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("My Middleware")
		next.ServeHTTP(w, r)
	})

}

var addr = flag.String("addr", ":3000", "hostname of server")

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)

	ba := bulma.BasicAuth("Auth", mux, bulma.User{
		"Toto": "1234",
		"Will": "0000",
	})

	handler := handlers.Logger(Middleware(ba))
	http.ListenAndServe(*addr, handler)

}
