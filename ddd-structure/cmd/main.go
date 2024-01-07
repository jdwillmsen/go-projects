package main

import (
	"github.com/google/uuid"
	"github.com/jdwillmsen/ddd-structure/domain/product"
	"github.com/jdwillmsen/ddd-structure/services/order"
	"github.com/jdwillmsen/ddd-structure/services/tavern"
)

func main() {
	products := productInventory()

	os, err := order.NewOrderService(
		//order.WithMongoCustomerRepository(context.Background(), "mongodb://admin:password@localhost:27017/?authSource=admin&authMechanism=SCRAM-SHA-1"),
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)

	if err != nil {
		panic(err)
	}

	tavern, err := tavern.NewTavern(
		tavern.WithOrderService(os),
	)

	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("Jake")
	if err != nil {
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy beverage", 1.99)
	if err != nil {
		panic(err)
	}

	peanuts, err := product.NewProduct("Peanuts", "Snacks", 0.99)

	if err != nil {
		panic(err)
	}

	wine, err := product.NewProduct("Wine", "Nasty drink", 0.99)
	if err != nil {
		panic(err)
	}

	return []product.Product{
		beer, peanuts, wine,
	}
}
