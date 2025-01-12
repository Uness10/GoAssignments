package repositories

import (
	"context"

	"bookstore.com/models"
)

type BookSaleStore interface {
	Create(ctx context.Context, book models.BookSale) (models.BookSale, error)
	Get(ctx context.Context, idx int) (models.BookSale, error)
	Delete(ctx context.Context, idx int) error
	Search(ctx context.Context, query models.SearchCriteria) ([]models.BookSale, error)
}
