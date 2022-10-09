package repository

import (
	"context"
	"github.com/baguss42/go-clean-arch/entity"
	"github.com/baguss42/go-clean-arch/entity/dto"
	_interface "github.com/baguss42/go-clean-arch/interface"
)

type ProductRepository struct {
	DB _interface.DatabaseInterface
}

func (r ProductRepository) All(ctx context.Context, option dto.ListOption) ([]entity.Product, error) {
	return nil, nil
}

func (r ProductRepository) Delete(ctx context.Context, id int64) error {
	return nil
}

func (r ProductRepository) Get(ctx context.Context, id int64) (entity.Product, error) {
	return entity.Product{}, nil
}

func (r ProductRepository) Insert(ctx context.Context, params dto.ProductParams) (entity.Product, error) {
	rows, err := r.DB.QueryRow(ctx, `INSERT INTO products (title, description, price, qty) VALUES ($1, $2, $3, $4) RETURNING id`,
		params.Title,
		params.Description,
		params.Price,
		params.Qty,
	)
	if err != nil {
		return entity.Product{}, err
	}

	var id int64
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return entity.Product{}, err
		}
	}

	return entity.Product{
		ID:          id,
		Title:       params.Title,
		Description: params.Description,
		Price:       params.Price,
		Qty:         params.Qty,
	}, nil
}

func (r ProductRepository) Update(ctx context.Context, id int64, params dto.ProductParams) (entity.Product, error) {
	return entity.Product{}, nil
}
