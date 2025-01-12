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

type BookSaleHandler struct {
	BookSaleService *services.BookSaleService
}

var (
	BookSaleInstance *BookSaleHandler
	BookSaleOnce     sync.Once
)

// NewBookSaleHandler initializes a singleton BookSaleInstance of BookSaleHandler.
func NewBookSaleHandler(BookSaleService *services.BookSaleService) *BookSaleHandler {
	BookSaleOnce.Do(func() {
		BookSaleInstance = &BookSaleHandler{BookSaleService: BookSaleService}
	})
	return BookSaleInstance
}

// CreateBookSale handles the creation of a new BookSale.
func (h *BookSaleHandler) CreateBookSale(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var BookSale models.BookSale
	err := json.NewDecoder(r.Body).Decode(&BookSale)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdBookSale, err := h.BookSaleService.CreateBookSale(BookSale)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdBookSale); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// GetBookSaleById retrieves a BookSale by its ID.
func (h *BookSaleHandler) GetBookSaleById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid BookSale ID", http.StatusBadRequest)
		return
	}

	BookSale, err := h.BookSaleService.GetBookSale(id)
	if err != nil {
		http.Error(w, "BookSale not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(BookSale); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// GetAllBookSales retrieves all BookSales.
func (h *BookSaleHandler) GetBookSalesByCriteria(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var query = models.SearchCriteria{
		Filters: make(map[string]interface{}),
	}
	err := json.NewDecoder(r.Body).Decode(&query.Filters)
	if err != nil {
		query = models.SearchCriteria{
			Filters: make(map[string]interface{}),
		}

	}

	// Call the service layer to search for BookSales based on criteria
	BookSales, err := h.BookSaleService.SearchBookSales(query)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the found BookSales
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(BookSales); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// DeleteBookSaleById deletes a BookSale by its ID.
func (h *BookSaleHandler) DeleteBookSaleById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid BookSale ID", http.StatusBadRequest)
		return
	}

	err = h.BookSaleService.DeleteBookSale(id)
	if err != nil {
		http.Error(w, "BookSale not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
