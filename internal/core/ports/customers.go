package ports

import "github.com/Edmartt/loyalty-system/internal/core/domain"

type Customers interface {
	ReadCustomer(dni string) (*domain.Customer, error)
	SaveCustomer(customer domain.Customer) (*domain.Customer, error)
}
