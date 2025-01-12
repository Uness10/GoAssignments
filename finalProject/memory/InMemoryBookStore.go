package memory

import (
	"context"
	"errors"
	"strings"
	"sync"

	"bookstore.com/models"
)

type InMemoryBookStore struct {
	mu     sync.Mutex
	books  map[int]models.Book
	nextID int
}

func NewInMemoryBookStore() *InMemoryBookStore {
	return &InMemoryBookStore{
		books:  make(map[int]models.Book),
		nextID: 1,
	}
}

func (s *InMemoryBookStore) Create(ctx context.Context, book models.Book) (models.Book, error) {
	select {
	case <-ctx.Done():
		return models.Book{}, ctx.Err()
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		book.ID = s.nextID
		s.books[s.nextID] = book
		s.nextID++
		return book, nil
	}
}

func (s *InMemoryBookStore) Get(ctx context.Context, id int) (models.Book, error) {
	select {
	case <-ctx.Done():
		return models.Book{}, ctx.Err()
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		book, exists := s.books[id]
		if !exists {
			return models.Book{}, errors.New("book not found")
		}
		return book, nil
	}
}

func (s *InMemoryBookStore) Update(ctx context.Context, book models.Book) (models.Book, error) {
	select {
	case <-ctx.Done():
		return models.Book{}, ctx.Err()
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		_, exists := s.books[book.ID]
		if !exists {
			return models.Book{}, errors.New("book not found")
		}
		s.books[book.ID] = book
		return book, nil
	}
}

func (s *InMemoryBookStore) Delete(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		_, exists := s.books[id]
		if !exists {
			return errors.New("book not found")
		}
		delete(s.books, id)
		return nil
	}
}

func (s *InMemoryBookStore) Search(ctx context.Context, query models.SearchCriteria) ([]models.Book, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		var results []models.Book

		for _, book := range s.books {
			match := true

			if title, exists := query.Filters["title"]; exists {
				if !strings.Contains(book.Title, title.(string)) {
					match = false
				}
			}

			if author, exists := query.Filters["author"]; exists {
				if !strings.Contains(book.Author.FirstName, author.(string)) {
					match = false
				}
			}

			if genre, exists := query.Filters["genre"]; exists {
				genreMatch := false
				for _, g := range book.Genres {
					if strings.Contains(g, genre.(string)) {
						genreMatch = true
						break
					}
				}
				if !genreMatch {
					match = false
				}
			}

			if price, exists := query.Filters["price"]; exists {
				if book.Price != price.(float64) {
					match = false
				}
			}
			if match {
				results = append(results, book)
			}
		}

		return results, nil
	}
}
