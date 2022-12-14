package controller

import (
	"context"
	controller "github.com/baguss42/go-clean-arch/controller/http_controller"
	"github.com/baguss42/go-clean-arch/service"
)

type Controller struct {
	Product controller.Product
}

func NewController(ctx context.Context, svc *service.Service) *Controller {
	return &Controller{
		Product: controller.Product{
			Ctx:     ctx,
			Service: svc.ProductService,
		},
	}
}
