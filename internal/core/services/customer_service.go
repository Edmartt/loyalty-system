package services

import (
	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/Edmartt/loyalty-system/internal/core/ports"
)

type CustomerService struct {
	CustomerRepository ports.Customers
}

func NewCustomerService(customerRepository ports.Customers) *CustomerService {
	return &CustomerService{
		CustomerRepository: customerRepository,
	}
}

func (s *CustomerService) CreateUser(customer domain.Customer) (*domain.Customer, error) {

	return s.CustomerRepository.SaveCustomer(customer)
}

func (s *CustomerService) GetUser(dni string) (*domain.Customer, error) {
	return s.CustomerRepository.ReadCustomer(dni)
}
