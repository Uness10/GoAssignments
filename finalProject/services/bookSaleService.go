package services

import (
	"context"

	"bookstore.com/models"
	"bookstore.com/repositories"
)

type BookSaleService struct {
	BookSaleRepo repositories.BookSaleStore
}

func NewBookSaleService(repo repositories.BookSaleStore) *BookSaleService {
	return &BookSaleService{BookSaleRepo: repo}
}

func (s *BookSaleService) CreateBookSale(ctx context.Context, BookSale models.BookSale) (models.BookSale, error) {
	return s.BookSaleRepo.Create(ctx, BookSale)
}

func (s *BookSaleService) GetBookSale(ctx context.Context, id int) (models.BookSale, error) {
	return s.BookSaleRepo.Get(ctx, id)
}

func (s *BookSaleService) DeleteBookSale(ctx context.Context, id int) error {
	return s.BookSaleRepo.Delete(ctx, id)
}

func (s *BookSaleService) SearchBookSales(ctx context.Context, query models.SearchCriteria) ([]models.BookSale, error) {
	return s.BookSaleRepo.Search(ctx, query)
}
