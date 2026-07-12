package service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/rajabhishekmaurya/ecommerce-microservices/payment-service/proto"
)

type PaymentServer struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *PaymentServer) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {

	fmt.Println("========== PAYMENT SERVICE ==========")
	fmt.Println("Order ID :", req.OrderId)
	fmt.Println("Amount   :", req.Amount)

	time.Sleep(500 * time.Millisecond)

	return &pb.PaymentResponse{
		Success:       true,
		TransactionId: fmt.Sprintf("TXN-%d", time.Now().Unix()),
	}, nil
}
