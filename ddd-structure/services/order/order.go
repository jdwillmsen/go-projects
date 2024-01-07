package order

import (
	"context"
	"github.com/google/uuid"
	"github.com/jdwillmsen/ddd-structure/domain/customer"
	"github.com/jdwillmsen/ddd-structure/domain/customer/memory"
	"github.com/jdwillmsen/ddd-structure/domain/customer/mongo"
	"github.com/jdwillmsen/ddd-structure/domain/product"
	prodmem "github.com/jdwillmsen/ddd-structure/domain/product/memory"
	"log"
)

type Configuration func(os *Service) error

type Service struct {
	customers customer.Repository
	products  product.Repository
}

func NewOrderService(cfgs ...Configuration) (*Service, error) {
	os := &Service{}
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.Repository) Configuration {
	return func(os *Service) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() Configuration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connStr string) Configuration {
	return func(os *Service) error {
		cr, err := mongo.New(ctx, connStr)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []product.Product) Configuration {
	return func(os *Service) error {
		pr := prodmem.New()

		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}

		os.products = pr
		return nil
	}
}

func (o *Service) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}
	var products []product.Product
	var total float64

	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)

		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	return total, nil
}

func (o *Service) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}

	err = o.customers.Add(c)
	if err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}
