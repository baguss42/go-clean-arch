package service

import (
	"github.com/baguss42/go-clean-arch/infrastructure/database"
)

type Service struct {
	ProductService *ProductService
}

func NewService(db *database.Database) *Service {
	return &Service{
		ProductService: NewProductService(db),
	}
}
