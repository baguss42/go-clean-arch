package grpc_controller

import (
	"context"
	"github.com/baguss42/go-clean-arch/app/grpc/pb"
	"github.com/baguss42/go-clean-arch/entity/dto"
	_interface "github.com/baguss42/go-clean-arch/interface"
)

type ProductServer struct {
	pb.UnimplementedProductServer
	Service _interface.ProductServiceInterface
}

func (server *ProductServer) Create(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	params := dto.ProductRequest{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		Qty:         int(req.GetQty()),
	}

	res, err := server.Service.Create(ctx, params)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		ID:          res.ID,
		Title:       res.Title,
		Description: res.Description,
		Price:       res.Price,
		Qty:         int32(res.Qty),
	}, nil
}
