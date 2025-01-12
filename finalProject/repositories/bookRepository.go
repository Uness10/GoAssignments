package repositories

import (
	"context"

	"bookstore.com/models"
)

type BookStore interface {
	Create(ctx context.Context, book models.Book) (models.Book, error)

	Get(ctx context.Context, idx int) (models.Book, error)

	Update(ctx context.Context, item models.Book) (models.Book, error)

	Delete(ctx context.Context, idx int) error

	Search(ctx context.Context, query models.SearchCriteria) ([]models.Book, error)
}
