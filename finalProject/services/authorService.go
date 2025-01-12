package services

import (
	"context"

	"bookstore.com/models"
	"bookstore.com/repositories"
)

type AuthorService struct {
	authorRepo repositories.AuthorStore
}

func NewAuthorService(repo repositories.AuthorStore) *AuthorService {
	return &AuthorService{authorRepo: repo}
}

func (s *AuthorService) CreateAuthor(ctx context.Context, author models.Author) (models.Author, error) {
	return s.authorRepo.Create(cts, author)
}

func (s *AuthorService) GetAuthor(ctx context.Context, id int) (models.Author, error) {
	return s.authorRepo.Get(ctx, id)
}

func (s *AuthorService) UpdateAuthor(ctx context.Context, author models.Author) (models.Author, error) {
	return s.authorRepo.Update(cts, author)
}

func (s *AuthorService) DeleteAuthor(ctx context.Context, id int) error {
	return s.authorRepo.Delete(ctx, id)
}

func (s *AuthorService) SearchAuthors(ctx context.Context, query models.SearchCriteria) ([]models.Author, error) {
	return s.authorRepo.Search(ctx, query)
}
