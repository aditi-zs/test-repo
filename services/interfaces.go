package services

import (
	"github.com/aditi-zs/test-repo/models"
	"gofr.dev/pkg/gofr"
)

type Product interface {
	GetAll(ctx *gofr.Context, f string) ([]*models.Product, error)
	Create(ctx *gofr.Context, product *models.Product) error
}
