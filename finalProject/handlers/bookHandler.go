package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
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

// NewBookHandler initializes a singleton instance of BookHandler.
func NewBookHandler(bookService *services.BookService) *BookHandler {
	once.Do(func() {
		instance = &BookHandler{bookService: bookService}
	})
	return instance
}

// CreateBook handles the creation of a new book.
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdBook, err := h.bookService.CreateBook(book)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdBook); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// GetBookById retrieves a book by its ID.
func (h *BookHandler) GetBookById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := h.bookService.GetBookByID(id)
	if err != nil {
		http.Error(w, "Book not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// GetBooksByCriteria retrieves all books.
func (h *BookHandler) GetBooksByCriteria(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var query = models.SearchCriteria{
		Filters: make(map[string]interface{}),
	}

	err := json.NewDecoder(r.Body).Decode(&query.Filters)
	if err != nil {
		query = models.SearchCriteria{
			Filters: make(map[string]interface{}),
		}

	}
	//fmt.Println(query)
	books, err := h.bookService.SearchBooks(query)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// UpdateBookById updates a book by its ID.
func (h *BookHandler) UpdateBookById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}
	book.ID = id

	updatedBook, err := h.bookService.UpdateBook(book)
	if err != nil {
		http.Error(w, "Book not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedBook); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// DeleteBookById deletes a book by its ID.
func (h *BookHandler) DeleteBookById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = h.bookService.DeleteBook(id)
	if err != nil {
		http.Error(w, "Book not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
