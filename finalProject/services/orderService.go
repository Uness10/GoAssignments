package services

import (
	"context"

	"bookstore.com/models"
	"bookstore.com/repositories"
)

type OrderService struct {
	orderRepo repositories.OrderStore
}

func NewOrderService(repo repositories.OrderStore) *OrderService {
	return &OrderService{orderRepo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, order models.Order) (models.Order, error) {
	return s.orderRepo.Create(ctx, order)
}

func (s *OrderService) GetOrder(ctx context.Context, id int) (models.Order, error) {
	return s.orderRepo.Get(ctx, id)
}

func (s *OrderService) UpdateOrder(ctx context.Context, order models.Order) (models.Order, error) {
	return s.orderRepo.Update(ctx, order)
}

func (s *OrderService) DeleteOrder(ctx context.Context, id int) error {
	return s.orderRepo.Delete(ctx, id)
}

func (s *OrderService) SearchOrders(ctx context.Context, query models.SearchCriteria) ([]models.Order, error) {
	return s.orderRepo.Search(ctx, query)
}
