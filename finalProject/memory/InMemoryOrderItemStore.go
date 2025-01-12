package memory

import (
	"context"
	"errors"
	"sync"

	"bookstore.com/models"
)

type InMemoryOrderItemStore struct {
	mu         sync.Mutex
	OrderItems map[int]models.OrderItem
	nextID     int
}

func NewInMemoryOrderItemStore() *InMemoryOrderItemStore {
	return &InMemoryOrderItemStore{
		OrderItems: make(map[int]models.OrderItem),
		nextID:     1,
	}
}

// Create adds a new order item to the store with context support.
func (s *InMemoryOrderItemStore) Create(ctx context.Context, orderItem models.OrderItem) (models.OrderItem, error) {
	select {
	case <-ctx.Done():
		return models.OrderItem{}, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		orderItem.ID = s.nextID
		s.OrderItems[s.nextID] = orderItem
		s.nextID++
		return orderItem, nil
	}
}

// Get retrieves an order item by its ID from the store with context support.
func (s *InMemoryOrderItemStore) Get(ctx context.Context, id int) (models.OrderItem, error) {
	select {
	case <-ctx.Done():
		return models.OrderItem{}, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		orderItem, exists := s.OrderItems[id]
		if !exists {
			return models.OrderItem{}, errors.New("OrderItem not found")
		}
		return orderItem, nil
	}
}

// Update modifies an existing order item in the store with context support.
func (s *InMemoryOrderItemStore) Update(ctx context.Context, orderItem models.OrderItem) (models.OrderItem, error) {
	select {
	case <-ctx.Done():
		return models.OrderItem{}, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		_, exists := s.OrderItems[orderItem.ID]
		if !exists {
			return models.OrderItem{}, errors.New("OrderItem not found")
		}
		s.OrderItems[orderItem.ID] = orderItem
		return orderItem, nil
	}
}

// Delete removes an order item from the store by its ID with context support.
func (s *InMemoryOrderItemStore) Delete(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		_, exists := s.OrderItems[id]
		if !exists {
			return errors.New("OrderItem not found")
		}
		delete(s.OrderItems, id)
		return nil
	}
}

// Search finds order items in the store based on search criteria with context support.
func (s *InMemoryOrderItemStore) Search(ctx context.Context, query models.SearchCriteria) ([]models.OrderItem, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		var results []models.OrderItem
		for _, orderItem := range s.OrderItems {
			results = append(results, orderItem)
		}
		return results, nil
	}
}
