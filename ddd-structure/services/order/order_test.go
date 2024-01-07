package order

import (
	"github.com/google/uuid"
	"github.com/jdwillmsen/ddd-structure/domain/customer"
	"github.com/jdwillmsen/ddd-structure/services/utilities"
	"testing"
)

func TestOrder_NewOrderService(t *testing.T) {
	products := utilities.InitProducts(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	cust, err := customer.NewCustomer("Jake")
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
