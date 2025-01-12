package memory

import (
	"context"
	"errors"
	"sync"

	"bookstore.com/models"
)

type InMemoryOrderStore struct {
	mu     sync.Mutex
	Orders map[int]models.Order
	nextID int
}

func NewInMemoryOrderStore() *InMemoryOrderStore {
	return &InMemoryOrderStore{
		Orders: make(map[int]models.Order),
		nextID: 1,
	}
}

// Create adds a new order to the store with context support.
func (s *InMemoryOrderStore) Create(ctx context.Context, order models.Order) (models.Order, error) {
	select {
	case <-ctx.Done():
		return models.Order{}, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		order.ID = s.nextID
		s.Orders[s.nextID] = order
		s.nextID++
		return order, nil
	}
}

// Get retrieves an order by its ID from the store with context support.
func (s *InMemoryOrderStore) Get(ctx context.Context, id int) (models.Order, error) {
	select {
	case <-ctx.Done():
		return models.Order{}, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		order, exists := s.Orders[id]
		if !exists {
			return models.Order{}, errors.New("Order not found")
		}
		return order, nil
	}
}

// Update modifies an existing order in the store with context support.
func (s *InMemoryOrderStore) Update(ctx context.Context, order models.Order) (models.Order, error) {
	select {
	case <-ctx.Done():
		return models.Order{}, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		_, exists := s.Orders[order.ID]
		if !exists {
			return models.Order{}, errors.New("Order not found")
		}
		s.Orders[order.ID] = order
		return order, nil
	}
}

// Delete removes an order from the store by its ID with context support.
func (s *InMemoryOrderStore) Delete(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		_, exists := s.Orders[id]
		if !exists {
			return errors.New("Order not found")
		}
		delete(s.Orders, id)
		return nil
	}
}

// Search finds orders in the store based on search criteria with context support.
func (s *InMemoryOrderStore) Search(ctx context.Context, query models.SearchCriteria) ([]models.Order, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		var results []models.Order
		for _, order := range s.Orders {
			results = append(results, order)
		}
		return results, nil
	}
}
