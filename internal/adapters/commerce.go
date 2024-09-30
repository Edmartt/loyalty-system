package adapters

import (
	"fmt"

	"github.com/Edmartt/loyalty-system/internal/core/domain"
	ports "github.com/Edmartt/loyalty-system/internal/core/ports/database"
)

type CommerceRepository struct {
	dbConnection ports.IConnection
}

func NewCommerceRepository(dbConnection ports.IConnection) *CommerceRepository {
	return &CommerceRepository{
		dbConnection: dbConnection,
	}
}

func (c *CommerceRepository) SaveCommerce(domain.Commerce) (*domain.Commerce, error) {
	dbConnection, err := c.dbConnection.GetConnection()

	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}

	_, err = dbConnection.NamedExec("INSERT INTO commerce (name, points_x_buy, value_x_point) VALUES(:name, :points_x_buy, :value_x_point)", &domain.Commerce{})

	if err != nil {
		return nil, fmt.Errorf("error saving data: %v", err)
	}

	return &domain.Commerce{}, nil
}
