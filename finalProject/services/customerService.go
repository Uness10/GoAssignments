package services

import (
	"context"

	"bookstore.com/models"
	"bookstore.com/repositories"
)

type CustomerService struct {
	customerRepo repositories.CustomerStore
}

func NewCustomerService(repo repositories.CustomerStore) *CustomerService {
	return &CustomerService{customerRepo: repo}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, customer models.Customer) (models.Customer, error) {
	return s.customerRepo.Create(ctx, customer)
}

func (s *CustomerService) GetCustomer(ctx context.Context, id int) (models.Customer, error) {
	return s.customerRepo.Get(ctx, id)
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, customer models.Customer) (models.Customer, error) {
	return s.customerRepo.Update(ctx, customer)
}

func (s *CustomerService) DeleteCustomer(ctx context.Context, id int) error {
	return s.customerRepo.Delete(ctx, id)
}

func (s *CustomerService) SearchCustomers(ctx context.Context, query models.SearchCriteria) ([]models.Customer, error) {
	return s.customerRepo.Search(ctx, query)
}
