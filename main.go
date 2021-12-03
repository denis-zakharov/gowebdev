package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "<h1>Welcome to my great site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch email me at "+
		"<a href=\"mailto:dizaharov@gmail.com\">dizaharov@gmail.com</a>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/bazinga":
		http.Error(w, "Bazinga!", http.StatusNotFound)
	default:
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		msg := http.StatusText(http.StatusNotFound)
		fmt.Fprintf(w,
			"%s<p>Path: %s</p><p>Raw (encoded) path:%s</p>",
			msg, r.URL.Path, r.URL.RawPath)
	}
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathHandler(w, r)
}

func main() {
	var router Router
	fmt.Println("Starting the web server on 3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
