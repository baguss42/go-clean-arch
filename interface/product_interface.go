package _interface

import (
	"context"
	"github.com/baguss42/go-clean-arch/entity"
	"github.com/baguss42/go-clean-arch/entity/dto"
)

type ProductServiceInterface interface {
	Create(ctx context.Context, request dto.ProductRequest) (response dto.ProductResponse, err error)
	Delete(ctx context.Context, id int64) error
	Detail(ctx context.Context, id int64) (response dto.ProductResponse, err error)
	List(ctx context.Context, request dto.ListOption) (response dto.ProductResponse, err error)
	Update(ctx context.Context, id int64, request dto.ProductRequest) (response dto.ProductResponse, err error)
}

type ProductRepositoryInterface interface {
	All(ctx context.Context, option dto.ListOption) ([]entity.Product, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (entity.Product, error)
	Insert(ctx context.Context, params dto.ProductParams) (entity.Product, error)
	Update(ctx context.Context, id int64, params dto.ProductParams) (entity.Product, error)
}
