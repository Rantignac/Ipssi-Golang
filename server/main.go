package main

import (
	"flag"
	"fmt"
	"html/template"
	"ipssi/handlers"
	"net/http"
	"strings"

	"github.com/nanoninja/bulma"
)

var filters = map[string]interface{}{
	"uppercase": strings.ToUpper,
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("home").Funcs(filters).ParseFiles("./tmpl/home.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, map[string]interface{}{
		"name":    "Gopher",
		"version": "1.10",
		"value":   "",
	})

	if err != nil {
		fmt.Println(err)
	}

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
