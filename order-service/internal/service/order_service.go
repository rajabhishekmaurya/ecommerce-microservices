package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/rajabhishekmaurya/ecommerce-microservices/order-service/internal/model"
	"github.com/rajabhishekmaurya/ecommerce-microservices/order-service/internal/repository"
)

type OrderService struct {
	repo          *repository.OrderRepository
	paymentClient *PaymentClient
}

func NewOrderService() (*OrderService, error) {

	paymentClient, err := NewPaymentClient()
	if err != nil {
		return nil, err
	}

	return &OrderService{
		repo:          repository.NewOrderRepository(),
		paymentClient: paymentClient,
	}, nil
}

func (s *OrderService) CreateOrder(req *model.CreateOrderRequest) (*model.CreateOrderResponse, error) {

	orderID := fmt.Sprintf("ORD-%d-%s", time.Now().Unix(), uuid.New().String()[:8])

	paymentResp, err := s.paymentClient.ProcessPayment(orderID, req.Amount)
	if err != nil {
		return nil, err
	}

	order := &model.Order{
		ID:            orderID,
		UserID:        req.UserID,
		ProductID:     req.ProductID,
		Amount:        req.Amount,
		Status:        "SUCCESS",
		TransactionID: paymentResp.TransactionId,
	}

	s.repo.Save(order)

	return &model.CreateOrderResponse{
		OrderID:       order.ID,
		TransactionID: order.TransactionID,
		Status:        order.Status,
	}, nil
}
