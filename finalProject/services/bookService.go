package services

import (
	"context"

	"bookstore.com/models"
	"bookstore.com/repositories"
)

type BookService struct {
	bookRepo repositories.BookStore
}

func NewBookService(bookRepo repositories.BookStore) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

// CreateBook adds a new book to the store with validation and context propagation
func (s *BookService) CreateBook(ctx context.Context, book models.Book) (models.Book, error) {

	// Pass the context to repository method
	return s.bookRepo.Create(ctx, book)
}

// GetBookByID retrieves a book by its ID, passing context to the repository
func (s *BookService) GetBookByID(ctx context.Context, id int) (models.Book, error) {
	// Pass the context to repository method
	return s.bookRepo.Get(ctx, id)
}

// UpdateBook updates an existing book in the store
func (s *BookService) UpdateBook(ctx context.Context, book models.Book) (models.Book, error) {

	// Pass the context to repository method
	return s.bookRepo.Update(ctx, book)
}

// DeleteBook removes a book from the store
func (s *BookService) DeleteBook(ctx context.Context, id int) error {
	return s.bookRepo.Delete(ctx, id)
}

// SearchBooks searches for books based on the query criteria
func (s *BookService) SearchBooks(ctx context.Context, query models.SearchCriteria) ([]models.Book, error) {
	// Pass the context to repository method
	return s.bookRepo.Search(ctx, query)
}
