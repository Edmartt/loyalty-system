package adapters

import (
	"fmt"

	"github.com/Edmartt/loyalty-system/internal/core/domain"
	ports "github.com/Edmartt/loyalty-system/internal/core/ports/database"
)

type BranchRepository struct {
	dbConnection ports.IConnection
}

func NewBranchRepository(dbConnection ports.IConnection) *BranchRepository {
	return &BranchRepository{
		dbConnection: dbConnection,
	}
}

func (b *BranchRepository) CreateBranch(branch domain.Branch) (*domain.Branch, error) {
	dbConnection, err := b.dbConnection.GetConnection()

	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}

	_, err = dbConnection.NamedExec("INSERT INTO branch (id, commerce_id, address) VALUES(:id, :commerce_id, :address)", branch)

	if err != nil {
		return nil, fmt.Errorf("error saving data: %v", err)
	}

	return &branch, nil
}
