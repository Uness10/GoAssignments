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

type CustomerHandler struct {
	CustomerService *services.CustomerService
}

var (
	CustomerInstance *CustomerHandler
	CustomerOnce     sync.Once
)

// NewCustomerHandler initializes a singleton CustomerInstance of CustomerHandler.
func NewCustomerHandler(CustomerService *services.CustomerService) *CustomerHandler {
	CustomerOnce.Do(func() {
		CustomerInstance = &CustomerHandler{CustomerService: CustomerService}
	})
	return CustomerInstance
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var Customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&Customer)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdCustomer, err := h.CustomerService.CreateCustomer(Customer)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdCustomer); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *CustomerHandler) GetCustomerById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid Customer ID", http.StatusBadRequest)
		return
	}

	Customer, err := h.CustomerService.GetCustomer(id)
	if err != nil {
		http.Error(w, "Customer not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Customer); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// GetAllCustomers retrieves all Customers.
func (h *CustomerHandler) GetCustomersByCriteria(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var query = models.SearchCriteria{
		Filters: make(map[string]interface{}),
	}
	err := json.NewDecoder(r.Body).Decode(&query.Filters)
	if err != nil {
		query = models.SearchCriteria{
			Filters: make(map[string]interface{}),
		}

	}
	// Call the service layer to search for Customers based on criteria
	Customers, err := h.CustomerService.SearchCustomers(query)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the found Customers
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Customers); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// UpdateCustomerById updates a Customer by its ID.
func (h *CustomerHandler) UpdateCustomerById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid Customer ID", http.StatusBadRequest)
		return
	}

	var Customer models.Customer
	err = json.NewDecoder(r.Body).Decode(&Customer)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}
	Customer.ID = id

	updatedCustomer, err := h.CustomerService.UpdateCustomer(Customer)
	if err != nil {
		http.Error(w, "Customer not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedCustomer); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// DeleteCustomerById deletes a Customer by its ID.
func (h *CustomerHandler) DeleteCustomerById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid Customer ID", http.StatusBadRequest)
		return
	}

	err = h.CustomerService.DeleteCustomer(id)
	if err != nil {
		http.Error(w, "Customer not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
