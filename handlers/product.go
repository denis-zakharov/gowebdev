package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/denis-zakharov/gowebdev/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(w, "Unable to marshall json", http.StatusInternalServerError)
	}
	w.Write(d)
}
