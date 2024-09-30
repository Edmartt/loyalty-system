package adapters

import (
	"fmt"

	"github.com/Edmartt/loyalty-system/internal/core/domain"
	ports "github.com/Edmartt/loyalty-system/internal/core/ports/database"
)

type CustomersRepository struct {
	dbConn ports.IConnection
}

func NewCustomersRepository(dbConn ports.IConnection) *CustomersRepository {
	return &CustomersRepository{
		dbConn: dbConn,
	}
}

func (c *CustomersRepository) ReadCustomer(dni string) (*domain.Customer, error) {
	dbConnection, err := c.dbConn.GetConnection()

	var customer domain.Customer

	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}

	err = dbConnection.Get(&customer, "SELECT user_dni, name, last_name, phone, email, points_collected, cashback_collected FROM customer WHERE dni=?", dni)
	if err != nil {
		return nil, fmt.Errorf("error fetching user data: %v", err)
	}

	return &customer, nil
}

func (c *CustomersRepository) SaveCustomer(customer domain.Customer) (*domain.Customer, error) {
	dbConnection, err := c.dbConn.GetConnection()

	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}

	_, err = dbConnection.NamedExec("INSERT INTO customer(user_dni, name, last_name, phone, email, points_collected, cashback_collected) VALUES(:user_dni, :name, :last_name, :phone, :email, :points_collected, :cashback_collected)", &customer)

	if err != nil {
		return nil, fmt.Errorf("error inserting data: %v", err)
	}

	return &customer, nil
}
