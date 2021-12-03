package handlers

import (
	"log"
	"net/http"

	"github.com/denis-zakharov/gowebdev/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}
	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (*Products) getProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshall json", http.StatusInternalServerError)
	}
}
