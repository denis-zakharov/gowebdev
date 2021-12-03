package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{}
	p.Name = "Latte"
	p.Price = 4.9
	p.SKU = "blah-one-two"
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
