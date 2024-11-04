package products

import (
	"github.com/aditi-zs/test-repo/models"
	"github.com/aditi-zs/test-repo/services"
	"github.com/aditi-zs/test-repo/stores"
	"gofr.dev/pkg/gofr"
)

type service struct {
	productStore stores.Product
}

func New(productStore stores.Product) services.Product {
	return &service{productStore: productStore}
}

func (s *service) GetAll(ctx *gofr.Context) ([]*models.Product, error) {
	products, err := s.productStore.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Create(ctx *gofr.Context, product *models.Product) error {
	err := s.productStore.Create(ctx, product)
	if err != nil {
		return err
	}

	return nil
}