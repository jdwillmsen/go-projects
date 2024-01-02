package customer

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jdwillmsen/ddd-example/aggregate"
)

var (
	ErrCustomerNotFound    = errors.New(" the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer      = errors.New("failed ot update the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
