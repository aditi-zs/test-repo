package handlers

import "gofr.dev/pkg/gofr"

type Product interface {
	Get(ctx *gofr.Context) (interface{}, error)
}
