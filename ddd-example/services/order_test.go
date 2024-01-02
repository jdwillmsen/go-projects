package services

import (
	"github.com/google/uuid"
	"github.com/jdwillmsen/ddd-example/aggregate"
	"testing"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy beverage", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	peanuts, err := aggregate.NewProduct("Peanuts", "Snacks", 0.99)

	if err != nil {
		t.Fatal(err)
	}

	wine, err := aggregate.NewProduct("Wine", "Nasty drink", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{
		beer, peanuts, wine,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("Jake")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)

	if err != nil {
		t.Error(err)
	}
}
