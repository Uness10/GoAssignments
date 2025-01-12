package repositories

import (
	"context"

	"bookstore.com/models"
)

type OrderStore interface {
	Create(ctx context.Context, Order models.Order) (models.Order, error)
	Get(ctx context.Context, idx int) (models.Order, error)
	Update(ctx context.Context, item models.Order) (models.Order, error)
	Delete(ctx context.Context, idx int) error
	Search(ctx context.Context, query models.SearchCriteria) ([]models.Order, error)
}
