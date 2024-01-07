package utilities

import (
	"github.com/jdwillmsen/ddd-structure/domain/product"
	"testing"
)

func InitProducts(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy beverage", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	peanuts, err := product.NewProduct("Peanuts", "Snacks", 0.99)

	if err != nil {
		t.Fatal(err)
	}

	wine, err := product.NewProduct("Wine", "Nasty drink", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{
		beer, peanuts, wine,
	}
}
