package handler

import (
	"context"

	pb "github.com/rajabhishekmaurya/ecommerce-microservices/common/proto/payment"
	"github.com/rajabhishekmaurya/ecommerce-microservices/payment-service/internal/service"
)

type PaymentHandler struct {
	pb.UnimplementedPaymentServiceServer

	paymentService *service.PaymentService
}

func NewPaymentHandler(ps *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: ps,
	}
}

func (h *PaymentHandler) ProcessPayment(
	ctx context.Context,
	req *pb.PaymentRequest,
) (*pb.PaymentResponse, error) {
	return h.paymentService.ProcessPayment(ctx, req)
}
