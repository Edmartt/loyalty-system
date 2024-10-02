package tests

import (
	"testing"

	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/Edmartt/loyalty-system/internal/core/services"
	"github.com/Edmartt/loyalty-system/tests/mocks"
	"github.com/google/uuid"
)

func TestCustomerCreateUser(t *testing.T) {
	customerRepo := mocks.MockCustomerRepository{}

	customerService := services.NewCustomerService(&customerRepo)

	myTestUser := domain.Customer{
		ID:       uuid.NewString(),
		UserDNI:  "1000765123",
		Name:     "myTestUser",
		LastName: "serviceName",
		Phone:    "555-5555",
		Email:    "email@email.com",
	}

	userCreated, err := customerService.CreateUser(myTestUser)

	if err != nil {
		t.Errorf("error creating use: expected %v and got %v", userCreated, err)
	}
}

func TestCustomerReadUser(t *testing.T) {
	customerRepo := mocks.MockCustomerRepository{}

	customerService := services.NewCustomerService(&customerRepo)

	userResult, err := customerService.GetUser("1000765123")

	if err != nil {
		t.Errorf("error fetching user: expected %v and got %v", userResult, err)
	}
}
