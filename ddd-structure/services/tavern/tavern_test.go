package tavern

import (
	"context"
	"github.com/google/uuid"
	"github.com/jdwillmsen/ddd-structure/services/order"
	"github.com/jdwillmsen/ddd-structure/services/utilities"
	"testing"
)

func Test_Tavern(t *testing.T) {
	products := utilities.InitProducts(t)

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(context.Background(), "mongodb://admin:password@localhost:27017/?authSource=admin&authMechanism=SCRAM-SHA-1"),
		order.WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	uid, err := os.AddCustomer("Jake")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)

	if err != nil {
		t.Fatal(err)
	}
}
