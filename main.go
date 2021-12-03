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
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Not found</h1>")
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("Starting the web server on 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
