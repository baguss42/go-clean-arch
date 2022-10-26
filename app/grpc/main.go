package grpc

import (
	"fmt"
	"github.com/baguss42/go-clean-arch/app"
	"github.com/baguss42/go-clean-arch/app/grpc/pb"
	"github.com/baguss42/go-clean-arch/controller/grpc_controller"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Boot(engine *app.Engine) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", engine.Environment.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterServer(s, engine)

	log.Printf("GRPC starting at %d", engine.Environment.GRPCPort)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", engine.Environment.GRPCPort)
	}
}

func RegisterServer(s *grpc.Server, engine *app.Engine) {
	pb.RegisterProductServer(s, &grpc_controller.ProductServer{
		Service: engine.Service.ProductService,
	})
}
