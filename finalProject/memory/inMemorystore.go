package memory

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type InMemoryStore struct {
	BookStore      InMemoryBookStore
	AuthorStore    InMemoryAuthorStore
	CustomerStore  InMemoryCustomerStore
	OrderStore     InMemoryOrderStore
	OrderItemStore InMemoryOrderItemStore
}

var (
	instance *InMemoryStore
	once     sync.Once
)

func NewInMemoryStore() (*InMemoryStore, error) {
	var err error
	once.Do(func() {
		instance, err = LoadData()
	})

	if err != nil {
		log.Fatalf("Error loading data: %v", err)
		return nil, err
	}

	return instance, nil
}

func LoadData() (*InMemoryStore, error) {
	data, err := os.ReadFile("database.json")
	if err != nil {
		return nil, err
	}

	store := &InMemoryStore{}
	err = json.Unmarshal(data, &store)
	if err != nil {
		return nil, err
	}
	return store, nil
}
