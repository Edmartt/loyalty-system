package mocks

import (
	"fmt"

	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/google/uuid"
)

type MockCustomerRepository struct {
}

func (m *MockCustomerRepository) SaveCustomer(customer domain.Customer) (*domain.Customer, error) {

	if customer.ID == "" || customer.UserDNI == "" || customer.Name == "" || customer.LastName == "" || customer.Phone == "" || customer.Email == "" {
		return nil, fmt.Errorf("missing requiered information for user creation")
	}

	return &customer, nil
}

func (m *MockCustomerRepository) ReadCustomer(id string) (*domain.Customer, error) {

	return &domain.Customer{
		ID:       uuid.NewString(),
		UserDNI:  "1000765123",
		Name:     "my username",
		LastName: "my lastname",
		Email:    "email@email.com",
		Phone:    "555-5555",
	}, nil
}
