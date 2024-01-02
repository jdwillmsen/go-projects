package memory

import (
	"github.com/google/uuid"
	"github.com/jdwillmsen/ddd-example/aggregate"
	"github.com/jdwillmsen/ddd-example/domain/product"
	"sync"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, p := range mpr.products {
		products = append(products, p)
	}
	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if p, ok := mpr.products[id]; ok {
		return p, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

func (mpr *MemoryProductRepository) Add(newProd aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newProd.GetID()]; ok {
		return product.ErrProductAlreadyExists
	}

	mpr.products[newProd.GetID()] = newProd

	return nil
}

func (mpr *MemoryProductRepository) Update(update aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[update.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[update.GetID()] = update
	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(mpr.products, id)

	return nil
}
