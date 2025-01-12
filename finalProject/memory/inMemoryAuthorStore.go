package memory

import (
	"context"
	"errors"
	"strings"
	"sync"

	"bookstore.com/models"
)

type InMemoryAuthorStore struct {
	mu      sync.Mutex
	authors map[int]models.Author
	nextID  int
}

func NewInMemoryAuthorStore() *InMemoryAuthorStore {
	return &InMemoryAuthorStore{
		authors: make(map[int]models.Author),
		nextID:  1,
	}
}

// Create adds a new author to the store with context support.
func (s *InMemoryAuthorStore) Create(ctx context.Context, author models.Author) (models.Author, error) {
	select {
	case <-ctx.Done():
		return models.Author{}, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		author.ID = s.nextID
		s.authors[s.nextID] = author
		s.nextID++
		return author, nil
	}
}

// Get retrieves an author by its ID from the store with context support.
func (s *InMemoryAuthorStore) Get(ctx context.Context, id int) (models.Author, error) {
	select {
	case <-ctx.Done():
		return models.Author{}, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		author, exists := s.authors[id]
		if !exists {
			return models.Author{}, errors.New("Author not found")
		}
		return author, nil
	}
}

// Update modifies an existing author in the store with context support.
func (s *InMemoryAuthorStore) Update(ctx context.Context, author models.Author) (models.Author, error) {
	select {
	case <-ctx.Done():
		return models.Author{}, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		_, exists := s.authors[author.ID]
		if !exists {
			return models.Author{}, errors.New("Author not found")
		}
		s.authors[author.ID] = author
		return author, nil
	}
}

// Delete removes an author from the store by its ID with context support.
func (s *InMemoryAuthorStore) Delete(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		_, exists := s.authors[id]
		if !exists {
			return errors.New("Author not found")
		}
		delete(s.authors, id)
		return nil
	}
}

// Search finds authors in the store based on search criteria with context support.
func (s *InMemoryAuthorStore) Search(ctx context.Context, query models.SearchCriteria) ([]models.Author, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		var results []models.Author
		for _, author := range s.authors {
			match := true

			if firstName, exists := query.Filters["firstName"]; exists {
				if !strings.Contains(author.FirstName, firstName.(string)) {
					match = false
				}
			}

			if lastName, exists := query.Filters["lastName"]; exists {
				if !strings.Contains(author.LastName, lastName.(string)) {
					match = false
				}
			}

			if name, exists := query.Filters["name"]; exists {
				fullName := author.FirstName + " " + author.LastName
				if !strings.Contains(fullName, name.(string)) {
					match = false
				}
			}

			if match {
				results = append(results, author)
			}
		}

		return results, nil
	}
}
