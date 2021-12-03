package handlers

import (
	"net/http"
	"strconv"

	"github.com/denis-zakharov/gowebdev/data"
	"github.com/gorilla/mux"
)

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
