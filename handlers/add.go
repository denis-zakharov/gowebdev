package handlers

import (
	"net/http"

	"github.com/denis-zakharov/gowebdev/data"
)

func (psh *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	psh.l.Println("Handle POST Products")
	p := r.Context().Value(KeyProduct{}).(*data.Product)
	// store
	psh.l.Printf("Product: %#v", p)
	data.AddProduct(p)
}
