package repositories

import (
	"context"

	"bookstore.com/models"
)

type CustomerStore interface {
	Create(ctx context.Context, Customer models.Customer) (models.Customer, error)
	Get(ctx context.Context, idx int) (models.Customer, error)
	Update(ctx context.Context, item models.Customer) (models.Customer, error)
	Delete(ctx context.Context, idx int) error
	Search(ctx context.Context, query models.SearchCriteria) ([]models.Customer, error)
}
