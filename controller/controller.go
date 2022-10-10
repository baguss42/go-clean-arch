package controller

import (
	"context"
	"github.com/baguss42/go-clean-arch/infrastructure/database"
	"github.com/baguss42/go-clean-arch/service"
)

type Controller struct {
	Product
}

func NewController(ctx context.Context, db *database.Database) *Controller {
	return &Controller{
		Product{
			Service: service.NewProductService(db),
			Ctx:     ctx,
		},
	}
}
