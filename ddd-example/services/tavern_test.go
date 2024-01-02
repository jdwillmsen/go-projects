package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/jdwillmsen/ddd-example/aggregate"
	"testing"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("Jake")
	if err != nil {
		t.Fatal(err)
	}

	if err = os.customers.Add(cust); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(cust.GetID(), order)

	if err != nil {
		t.Fatal(err)
	}
}
