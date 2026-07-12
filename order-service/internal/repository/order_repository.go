package repository

import (
	"sync"

	"github.com/rajabhishekmaurya/ecommerce-microservices/order-service/internal/model"
)

type OrderRepository struct {
	mu     sync.RWMutex
	orders map[string]*model.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]*model.Order),
	}
}

func (r *OrderRepository) Save(order *model.Order) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[order.ID] = order
}

func (r *OrderRepository) GetByID(id string) (*model.Order, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, ok := r.orders[id]
	return order, ok
}
