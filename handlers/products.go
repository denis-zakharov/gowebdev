package handlers

import (
	"context"
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
	p := r.Context().Value(KeyProduct{}).(*data.Product)
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

	p := r.Context().Value(KeyProduct{}).(*data.Product)

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

type KeyProduct struct{}

func (psh *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := &data.Product{}
		err := p.FromJSON(r.Body)
		if err != nil {
			psh.l.Println("[ERROR] desirializing product", err)
			http.Error(w, "Error reading product", http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, p)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
