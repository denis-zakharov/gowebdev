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

func (psh *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		psh.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		psh.addProduct(w, r)
		return
	}

	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (psh *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	psh.l.Println("Handle GET Products")
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshall json", http.StatusInternalServerError)
	}
}

func (psh *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	psh.l.Println("Handle POST Products")
	p := &data.Product{}
	err := p.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshall json", http.StatusBadRequest)
		return
	}
	// store
	psh.l.Printf("Product: %#v", p)
}
