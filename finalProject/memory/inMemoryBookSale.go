package memory

import (
	"context"
	"errors"
	"strings"
	"sync"

	"bookstore.com/models"
)

type InMemorybookSalestore struct {
	mu        sync.Mutex
	bookSales map[int]models.BookSale
	nextID    int
}

func NewInMemorybookSalestore() *InMemorybookSalestore {
	return &InMemorybookSalestore{
		bookSales: make(map[int]models.BookSale),
		nextID:    1,
	}
}

func (s *InMemorybookSalestore) Create(ctx context.Context, BookSale models.BookSale) (models.BookSale, error) {
	select {
	case <-ctx.Done():
		return models.BookSale{}, ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	BookSale.ID = s.nextID
	s.bookSales[s.nextID] = BookSale
	s.nextID++
	return BookSale, nil
}

func (s *InMemorybookSalestore) Get(ctx context.Context, id int) (models.BookSale, error) {
	select {
	case <-ctx.Done():
		return models.BookSale{}, ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	BookSale, exists := s.bookSales[id]
	if !exists {
		return models.BookSale{}, errors.New("BookSale not found")
	}
	return BookSale, nil
}

func (s *InMemorybookSalestore) Update(ctx context.Context, BookSale models.BookSale) (models.BookSale, error) {
	select {
	case <-ctx.Done():
		return models.BookSale{}, ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.bookSales[BookSale.ID]
	if !exists {
		return models.BookSale{}, errors.New("BookSale not found")
	}
	s.bookSales[BookSale.ID] = BookSale
	return BookSale, nil
}

func (s *InMemorybookSalestore) Delete(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.bookSales[id]
	if !exists {
		return errors.New("BookSale not found")
	}
	delete(s.bookSales, id)
	return nil
}

func (s *InMemorybookSalestore) Search(ctx context.Context, query models.SearchCriteria) ([]models.BookSale, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var results []models.BookSale

	for _, obj := range s.bookSales {
		match := true

		if title, exists := query.Filters["title"]; exists {
			if !strings.Contains(obj.Book.Title, title.(string)) {
				match = false
			}
		}

		if author, exists := query.Filters["author"]; exists {
			if !strings.Contains(obj.Book.Author.FirstName, author.(string)) {
				match = false
			}
		}

		if genre, exists := query.Filters["genre"]; exists {
			genreMatch := false
			for _, g := range obj.Book.Genres {
				if strings.Contains(g, genre.(string)) {
					genreMatch = true
					break
				}
			}
			if !genreMatch {
				match = false
			}
		}

		if quantity, exists := query.Filters["quantity"]; exists {
			if obj.Quantity != quantity {
				match = false
			}
		}
		if match {
			results = append(results, obj)
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}

	return results, nil
}
