package services

import (
	"context"

	"bookstore.com/models"
	"bookstore.com/repositories"
)

type OrderItemService struct {
	orderItemRepo repositories.OrderItemStore
}

func NewOrderItemService(repo repositories.OrderItemStore) *OrderItemService {
	return &OrderItemService{orderItemRepo: repo}
}

func (s *OrderItemService) CreateOrderItem(ctx context.Context, orderItem models.OrderItem) (models.OrderItem, error) {
	return s.orderItemRepo.Create(ctx, orderItem)
}

func (s *OrderItemService) GetOrderItem(ctx context.Context, id int) (models.OrderItem, error) {
	return s.orderItemRepo.Get(ctx, id)
}

func (s *OrderItemService) UpdateOrderItem(ctx context.Context, orderItem models.OrderItem) (models.OrderItem, error) {
	return s.orderItemRepo.Update(ctx, orderItem)
}

func (s *OrderItemService) DeleteOrderItem(ctx context.Context, id int) error {
	return s.orderItemRepo.Delete(ctx, id)
}

func (s *OrderItemService) SearchOrderItems(ctx context.Context, query models.SearchCriteria) ([]models.OrderItem, error) {
	return s.orderItemRepo.Search(ctx, query)
}
