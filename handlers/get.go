package handlers

import (
	"net/http"

	"github.com/denis-zakharov/gowebdev/data"
)

func (psh *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	psh.l.Println("Handle GET Products")
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshall json", http.StatusInternalServerError)
	}
}
