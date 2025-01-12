package repositories

import (
	"context"

	"bookstore.com/models"
)

type AuthorStore interface {
	Create(ctx context.Context, Author models.Author) (models.Author, error)
	Get(ctx context.Context, idx int) (models.Author, error)
	Update(ctx context.Context, item models.Author) (models.Author, error)
	Delete(ctx context.Context, idx int) error
	Search(ctx context.Context, query models.SearchCriteria) ([]models.Author, error)
}
