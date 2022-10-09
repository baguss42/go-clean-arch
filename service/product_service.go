package service

import (
	"context"
	"fmt"
	"github.com/baguss42/go-clean-arch/entity/dto"
	"github.com/baguss42/go-clean-arch/infrastructure/database"
	_interface "github.com/baguss42/go-clean-arch/interface"
	"github.com/baguss42/go-clean-arch/repository"
)

type ProductService struct {
	Repo _interface.ProductRepositoryInterface
}

func NewProductService(db *database.Database) *ProductService {
	return &ProductService{
		Repo: repository.ProductRepository{
			DB: db,
		},
	}
}

func (svc ProductService) Create(ctx context.Context, request dto.ProductRequest) (response dto.ProductResponse, err error) {
	product, err := svc.Repo.Insert(ctx, dto.ProductParams{
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Qty:         request.Qty,
	})
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.ProductResponse{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Qty:         product.Qty,
	}, nil
}

func (svc ProductService) Delete(ctx context.Context, id int64) error {
	return svc.Repo.Delete(ctx, id)
}

func (svc ProductService) Detail(ctx context.Context, id int64) (response dto.ProductResponse, err error) {
	_, err = svc.Repo.Get(ctx, id)
	return dto.ProductResponse{}, err
}

func (svc ProductService) List(ctx context.Context, request dto.ListOption) (response dto.ProductResponse, err error) {
	_, err = svc.Repo.All(ctx, request)
	return dto.ProductResponse{}, err
}

func (svc ProductService) Update(ctx context.Context, id int64, request dto.ProductRequest) (response dto.ProductResponse, err error) {
	_, err = svc.Repo.Update(ctx, id, dto.ProductParams{})
	return dto.ProductResponse{}, fmt.Errorf("unexpected error")
}
