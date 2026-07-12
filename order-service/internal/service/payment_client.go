package service

import (
	"context"
	"time"

	pb "github.com/rajabhishekmaurya/ecommerce-microservices/payment-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PaymentClient struct {
	client pb.PaymentServiceClient
	conn   *grpc.ClientConn
}

func NewPaymentClient() (*PaymentClient, error) {

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewPaymentServiceClient(conn)

	return &PaymentClient{
		client: client,
		conn:   conn,
	}, nil
}

func (p *PaymentClient) ProcessPayment(orderID string, amount float64) (*pb.PaymentResponse, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return p.client.ProcessPayment(ctx, &pb.PaymentRequest{
		OrderId: orderID,
		Amount:  amount,
	})
}

func (p *PaymentClient) Close() error {
	return p.conn.Close()
}
