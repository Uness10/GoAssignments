package memory

import (
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

func (s *InMemoryOrderItemStore) Create(OrderItem models.OrderItem) (models.OrderItem, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	OrderItem.ID = s.nextID
	s.OrderItems[s.nextID] = OrderItem
	s.nextID++
	return OrderItem, nil
}

func (s *InMemoryOrderItemStore) Get(id int) (models.OrderItem, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	OrderItem, exists := s.OrderItems[id]
	if !exists {
		return models.OrderItem{}, errors.New("OrderItem not found")
	}
	return OrderItem, nil
}

func (s *InMemoryOrderItemStore) Update(OrderItem models.OrderItem) (models.OrderItem, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.OrderItems[OrderItem.ID]
	if !exists {
		return models.OrderItem{}, errors.New("OrderItem not found")
	}
	s.OrderItems[OrderItem.ID] = OrderItem
	return OrderItem, nil
}

func (s *InMemoryOrderItemStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.OrderItems[id]
	if !exists {
		return errors.New("OrderItem not found")
	}
	delete(s.OrderItems, id)
	return nil
}

func (s *InMemoryOrderItemStore) Search(query models.SearchCriteria) ([]models.OrderItem, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var results []models.OrderItem
	for _, OrderItem := range s.OrderItems {
		results = append(results, OrderItem)
	}
	return results, nil
}
