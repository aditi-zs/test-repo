package products

import (
	"github.com/aditi-zs/test-repo/handlers"
	"github.com/aditi-zs/test-repo/models"
	"github.com/aditi-zs/test-repo/services"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"
	"gofr.dev/pkg/gofr/http/response"
)

type handler struct {
	productService services.Product
}

func New(productService services.Product) handlers.Product {
	return &handler{productService: productService}
}

func (h *handler) Get(ctx *gofr.Context) (interface{}, error) {
	f := ctx.Param("name")

	data, err := h.productService.GetAll(ctx, f)
	if err != nil {
		return nil, err
	}

	res := response.Raw{Data: data}

	return res, nil
}

func (h *handler) Create(ctx *gofr.Context) (interface{}, error) {
	var product models.Product

	if err := ctx.Bind(&product); err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"body"}}
	}

	if product.Name == "" {
		return nil, http.ErrorMissingParam{Params: []string{"name"}}
	}

	err := h.productService.Create(ctx, &product)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
