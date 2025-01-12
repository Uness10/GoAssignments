package memory

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type InMemoryStore struct {
	BookStore     InMemoryBookStore     `json:"books"`
	AuthorStore   InMemoryAuthorStore   `json:"authors"`
	CustomerStore InMemoryCustomerStore `json:"customers"`
	OrderStore    InMemoryOrderStore    `json:"orders"`
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

	// Ensure each store is initialized after loading
	initializeStores(instance)

	return instance, nil
}

func initializeStores(store *InMemoryStore) {
	// Initialize BookStore if it is not initialized
	if store.BookStore.Books == nil {
		store.BookStore = *NewInMemoryBookStore()
	}

	// Initialize AuthorStore if it is not initialized
	if store.AuthorStore.Authors == nil {
		store.AuthorStore = *NewInMemoryAuthorStore()
	}

	// Initialize CustomerStore if it is not initialized
	if store.CustomerStore.Customers == nil {
		store.CustomerStore = *NewInMemoryCustomerStore()
	}

	// Initialize OrderStore if it is not initialized
	if store.OrderStore.Orders == nil {
		store.OrderStore = *NewInMemoryOrderStore()
	}

}

func LoadData() (*InMemoryStore, error) {
	data, err := os.ReadFile("database.json")
	if err != nil {
		return nil, err
	}

	store := &InMemoryStore{}
	err = json.Unmarshal(data, store)
	if err != nil {
		return nil, err
	}

	return store, nil
}
