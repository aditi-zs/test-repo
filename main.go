package main

import (
	productHandler "github.com/aditi-zs/test-repo/handlers/products"
	productService "github.com/aditi-zs/test-repo/services/products"
	"github.com/aditi-zs/test-repo/stores/products"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	//app.Migrate(migrations.All())

	productStore := products.New()
	productSvc := productService.New(productStore)
	productHandlers := productHandler.New(productSvc)

	app.GET("/product", productHandlers.Get)

	app.Run()
}
