package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"bookstore.com/models"
	"bookstore.com/services"
	"github.com/julienschmidt/httprouter"
)

type BookHandler struct {
	bookService *services.BookService
}

var (
	instance *BookHandler
	once     sync.Once
)

func NewBookHandler(bookService *services.BookService) *BookHandler {
	once.Do(func() {
		instance = &BookHandler{bookService: bookService}
	})
	return instance
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdBook, err := h.bookService.CreateBook(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}
