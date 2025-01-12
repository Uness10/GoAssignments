package repositories

import (
	"context"

	"bookstore.com/models"
)

type OrderItemStore interface {
	Create(ctx context.Context, OrderItem models.OrderItem) (models.OrderItem, error)
	Get(ctx context.Context, idx int) (models.OrderItem, error)
	Update(ctx context.Context, item models.OrderItem) (models.OrderItem, error)
	Delete(ctx context.Context, idx int) error
	Search(ctx context.Context, query models.SearchCriteria) ([]models.OrderItem, error)
}
