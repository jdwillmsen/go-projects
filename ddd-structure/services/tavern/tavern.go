package tavern

import (
	"github.com/google/uuid"
	"github.com/jdwillmsen/ddd-structure/services/order"
	"log"
)

type Configuration func(os *Tavern) error

type Tavern struct {
	OrderService   *order.Service
	BillingService interface{}
}

func NewTavern(cfgs ...Configuration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		if err := cfg(t); err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithOrderService(os *order.Service) Configuration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	log.Printf("\nBill the customer: %0.0f\n", price)

	return nil
}
