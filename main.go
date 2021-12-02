package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "<h1>Welcome to my great site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the web server on 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
