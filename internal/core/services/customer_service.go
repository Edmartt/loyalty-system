package services

import (
	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/Edmartt/loyalty-system/internal/core/ports"
)

type Service struct {
	customerRepository ports.Customers
}

func NewCustomerService(customerRepository ports.Customers) *Service {
	return &Service{
		customerRepository: customerRepository,
	}
}

func (s *Service) CreateUser(customer domain.Customer) (*domain.Customer, error) {

	return s.customerRepository.SaveCustomer(customer)
}

func (s *Service) GetUser(dni string) (*domain.Customer, error) {
	return s.customerRepository.ReadCustomer(dni)
}
