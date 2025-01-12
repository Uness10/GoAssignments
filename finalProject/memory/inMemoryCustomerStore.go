package memory

import (
	"context"
	"errors"
	"sync"

	"bookstore.com/models"
)

type InMemoryCustomerStore struct {
	mu        sync.Mutex
	Customers map[int]models.Customer
	nextID    int
}

func NewInMemoryCustomerStore() *InMemoryCustomerStore {
	return &InMemoryCustomerStore{
		Customers: make(map[int]models.Customer),
		nextID:    1,
	}
}

// Create adds a new customer to the store with context support.
func (s *InMemoryCustomerStore) Create(ctx context.Context, customer models.Customer) (models.Customer, error) {
	select {
	case <-ctx.Done():
		return models.Customer{}, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		customer.ID = s.nextID
		s.Customers[s.nextID] = customer
		s.nextID++
		return customer, nil
	}
}

// Get retrieves a customer by its ID from the store with context support.
func (s *InMemoryCustomerStore) Get(ctx context.Context, id int) (models.Customer, error) {
	select {
	case <-ctx.Done():
		return models.Customer{}, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		customer, exists := s.Customers[id]
		if !exists {
			return models.Customer{}, errors.New("Customer not found")
		}
		return customer, nil
	}
}

// Update modifies an existing customer in the store with context support.
func (s *InMemoryCustomerStore) Update(ctx context.Context, customer models.Customer) (models.Customer, error) {
	select {
	case <-ctx.Done():
		return models.Customer{}, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		_, exists := s.Customers[customer.ID]
		if !exists {
			return models.Customer{}, errors.New("Customer not found")
		}
		s.Customers[customer.ID] = customer
		return customer, nil
	}
}

// Delete removes a customer from the store by its ID with context support.
func (s *InMemoryCustomerStore) Delete(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		_, exists := s.Customers[id]
		if !exists {
			return errors.New("Customer not found")
		}
		delete(s.Customers, id)
		return nil
	}
}

// Search finds customers in the store based on search criteria with context support.
func (s *InMemoryCustomerStore) Search(ctx context.Context, query models.SearchCriteria) ([]models.Customer, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err() // Handle cancellation or timeout
	default:
		s.mu.Lock()
		defer s.mu.Unlock()

		var results []models.Customer
		for _, customer := range s.Customers {
			results = append(results, customer)
		}
		return results, nil
	}
}
