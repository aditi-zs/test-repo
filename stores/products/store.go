package products

import (
	"github.com/aditi-zs/test-repo/models"
	"github.com/aditi-zs/test-repo/stores"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/datasource"
)

type store struct{}

func New() stores.Product {
	return &store{}
}

func (s *store) GetAll(ctx *gofr.Context) ([]*models.Product, error) {
	var products []*models.Product

	query := getAllQuery

	rows, err := ctx.SQL.QueryContext(ctx, query)
	if err != nil {
		return nil, datasource.ErrorDB{
			Err:     err,
			Message: "Query Execution Failed",
		}
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product

		err = rows.Scan(&product.ID, &product.Name)
		if err != nil {
			return nil, datasource.ErrorDB{
				Err:     err,
				Message: "DB Scan Error",
			}
		}

		products = append(products, &product)
	}

	return products, nil
}

func (s *store) Create(ctx *gofr.Context, product *models.Product) error {
	_, err := ctx.SQL.ExecContext(ctx, createQuery, product.Name)
	if err != nil {
		return datasource.ErrorDB{
			Err:     err,
			Message: "Query Execution Failed",
		}
	}

	return nil
}
