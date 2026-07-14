package service

import (
	"context"
	"fmt"
	"time"

	"github.com/rajabhishekmaurya/ecommerce-microservices/payment-service/internal/model"
	pb "github.com/rajabhishekmaurya/ecommerce-microservices/payment-service/proto"
)

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {

	fmt.Println("========== PAYMENT SERVICE ==========")
	fmt.Println("Order ID :", req.OrderId)
	fmt.Println("Amount   :", req.Amount)

	// Generate transaction ID once
	txnID := fmt.Sprintf("TXN-%d", time.Now().Unix())

	producer := NewKafkaProducer()
	defer producer.Close()

	event := &model.PaymentEvent{
		OrderID:       req.OrderId,
		TransactionID: txnID,
		Amount:        req.Amount,
		Status:        "SUCCESS",
	}

	if err := producer.Publish(event); err != nil {
		return nil, err
	}

	time.Sleep(500 * time.Millisecond)

	return &pb.PaymentResponse{
		Success:       true,
		TransactionId: txnID,
	}, nil
}
