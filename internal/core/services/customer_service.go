package services

import (
	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/Edmartt/loyalty-system/internal/core/ports"
)

type CustomerService struct {
	customerRepository ports.Customers
}

func NewCustomerService(customerRepository ports.Customers) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
	}
}

func (s *CustomerService) CreateUser(customer domain.Customer) (*domain.Customer, error) {

	return s.customerRepository.SaveCustomer(customer)
}

func (s *CustomerService) GetUser(dni string) (*domain.Customer, error) {
	return s.customerRepository.ReadCustomer(dni)
}
