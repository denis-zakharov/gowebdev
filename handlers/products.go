package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPut {
		// expect the id in the URI
		re := regexp.MustCompile(`/([0-9]+)`)
		g := re.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, _ := strconv.Atoi(idString)

		psh.l.Println("Trying to update product id", id)
		psh.updateProduct(id, w, r)
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
	data.AddProduct(p)
}

func (psh *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	psh.l.Println("Handle PUT Products")
	p := &data.Product{}
	err := p.FromJSON(r.Body)
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
