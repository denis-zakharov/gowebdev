package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye World!")
	fmt.Fprint(w, "Bye!")
}
