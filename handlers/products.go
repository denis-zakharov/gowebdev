package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/denis-zakharov/gowebdev/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (psh *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	psh.l.Println("Handle GET Products")
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshall json", http.StatusInternalServerError)
	}
}

func (psh *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	psh.l.Println("Handle POST Products")
	p := &data.Product{}
	err := p.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshall json", http.StatusBadRequest)
		return
	}
	// store
	psh.l.Printf("Product: %#v", p)
	data.AddProduct(p)
}

func (psh *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	psh.l.Println("Handle PUT Products")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p := &data.Product{}
	err = p.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshall json", http.StatusBadRequest)
		return
	}

	// update
	psh.l.Printf("Updating a Product: %#v", p)
	err = data.UpdateProduct(id, p)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product update failed", http.StatusInternalServerError)
		return
	}
}
