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

type AuthorHandler struct {
	AuthorService *services.AuthorService
}

var (
	AuthorInstance *AuthorHandler
	AuthorOnce     sync.Once
)

// NewAuthorHandler initializes a singleton AuthorInstance of AuthorHandler.
func NewAuthorHandler(AuthorService *services.AuthorService) *AuthorHandler {
	AuthorOnce.Do(func() {
		AuthorInstance = &AuthorHandler{AuthorService: AuthorService}
	})
	return AuthorInstance
}

// CreateAuthor handles the creation of a new Author.
func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var Author models.Author
	err := json.NewDecoder(r.Body).Decode(&Author)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdAuthor, err := h.AuthorService.CreateAuthor(Author)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdAuthor); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// GetAuthorById retrieves a Author by its ID.
func (h *AuthorHandler) GetAuthorById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid Author ID", http.StatusBadRequest)
		return
	}

	Author, err := h.AuthorService.GetAuthor(id)
	if err != nil {
		http.Error(w, "Author not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Author); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// GetAllAuthors retrieves all Authors.
func (h *AuthorHandler) GetAuthorsByCriteria(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var query = models.SearchCriteria{
		Filters: make(map[string]interface{}),
	}
	err := json.NewDecoder(r.Body).Decode(&query.Filters)
	if err != nil {
		query = models.SearchCriteria{
			Filters: make(map[string]interface{}),
		}

	}

	// Call the service layer to search for Authors based on criteria
	Authors, err := h.AuthorService.SearchAuthors(query)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the found Authors
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Authors); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// UpdateAuthorById updates a Author by its ID.
func (h *AuthorHandler) UpdateAuthorById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid Author ID", http.StatusBadRequest)
		return
	}

	var Author models.Author
	err = json.NewDecoder(r.Body).Decode(&Author)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}
	Author.ID = id

	updatedAuthor, err := h.AuthorService.UpdateAuthor(Author)
	if err != nil {
		http.Error(w, "Author not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedAuthor); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// DeleteAuthorById deletes a Author by its ID.
func (h *AuthorHandler) DeleteAuthorById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid Author ID", http.StatusBadRequest)
		return
	}

	err = h.AuthorService.DeleteAuthor(id)
	if err != nil {
		http.Error(w, "Author not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
