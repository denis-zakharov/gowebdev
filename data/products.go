package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"` // "-" to omit in serialization
}

type Products []*Product

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
		matches := re.FindAllString(fl.Field().String(), -1)
		return len(matches) == 1
	})
	return validate.Struct(p)
}

func (ps *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ps)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}
	return -1, ErrProductNotFound
}

func getNextID() int {
	last := productList[len(productList)-1]
	return last.ID + 1
}

var productList = Products{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Forthy milky cofee",
		Price:       2.45,
		SKU:         "abcd323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
