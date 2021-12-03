package handlers

import (
	"net/http"
	"strconv"

	"github.com/denis-zakharov/gowebdev/data"
	"github.com/gorilla/mux"
)

func (psh *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	psh.l.Println("Handle DELETE Products")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	// delete
	psh.l.Printf("Deleting a Product: %d", id)
	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product delete failed", http.StatusInternalServerError)
		return
	}
}
