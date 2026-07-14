package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/rajabhishekmaurya/ecommerce-microservices/payment-service/internal/config"
	"github.com/rajabhishekmaurya/ecommerce-microservices/payment-service/internal/handler"
	"github.com/rajabhishekmaurya/ecommerce-microservices/payment-service/internal/service"

	pb "github.com/rajabhishekmaurya/ecommerce-microservices/payment-service/proto"
)

type Server struct {
	cfg *config.Config
}

func New() *Server {
	return &Server{
		cfg: config.Load(),
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", ":"+s.cfg.Port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	paymentService := service.NewPaymentService()
	paymentHandler := handler.NewPaymentHandler(paymentService)

	pb.RegisterPaymentServiceServer(
		grpcServer,
		paymentHandler,
	)

	log.Printf("Payment Service listening on :%s", s.cfg.Port)

	return grpcServer.Serve(lis)
}
