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

type OrderHandler struct {
	OrderService *services.OrderService
}

var (
	OrderInstance *OrderHandler
	OrderOnce     sync.Once
)

// NewOrderHandler initializes a singleton OrderInstance of OrderHandler.
func NewOrderHandler(OrderService *services.OrderService) *OrderHandler {
	OrderOnce.Do(func() {
		OrderInstance = &OrderHandler{OrderService: OrderService}
	})
	return OrderInstance
}

// CreateOrder handles the creation of a new Order.
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var Order models.Order
	err := json.NewDecoder(r.Body).Decode(&Order)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdOrder, err := h.OrderService.CreateOrder(Order)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdOrder); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// GetOrderById retrieves a Order by its ID.
func (h *OrderHandler) GetOrderById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid Order ID", http.StatusBadRequest)
		return
	}

	Order, err := h.OrderService.GetOrder(id)
	if err != nil {
		http.Error(w, "Order not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Order); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// GetAllOrders retrieves all Orders.
func (h *OrderHandler) GetOrdersByCriteria(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var query = models.SearchCriteria{
		Filters: make(map[string]interface{}),
	}
	err := json.NewDecoder(r.Body).Decode(&query.Filters)
	if err != nil {
		query = models.SearchCriteria{
			Filters: make(map[string]interface{}),
		}

	}

	// Call the service layer to search for Orders based on criteria
	Orders, err := h.OrderService.SearchOrders(query)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the found Orders
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Orders); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// UpdateOrderById updates a Order by its ID.
func (h *OrderHandler) UpdateOrderById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid Order ID", http.StatusBadRequest)
		return
	}

	var Order models.Order
	err = json.NewDecoder(r.Body).Decode(&Order)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}
	Order.ID = id

	updatedOrder, err := h.OrderService.UpdateOrder(Order)
	if err != nil {
		http.Error(w, "Order not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedOrder); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}
}

// DeleteOrderById deletes a Order by its ID.
func (h *OrderHandler) DeleteOrderById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid Order ID", http.StatusBadRequest)
		return
	}

	err = h.OrderService.DeleteOrder(id)
	if err != nil {
		http.Error(w, "Order not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
