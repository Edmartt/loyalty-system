package adapters

import (
	"fmt"

	"github.com/Edmartt/loyalty-system/internal/adapters/http/dto"
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

func (b *BranchRepository) ReadBranchCampaign(id string) ([]dto.BranchCampaign, error) {

	dbConnection, err := b.dbConnection.GetConnection()

	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %v", err)
	}

	var result []dto.BranchCampaign

	err = dbConnection.Select(&result, "SELECT c.id AS campaign_id, c.campaign_name, c.campaign_multiplier, c.campaign_percent_bonus, c.start_date, c.end_date, b.id AS branch_id, b.address, co.id AS commerce_id, co.name AS commerce_name, co.points_x_buy, co.value_x_point FROM campaign c JOIN branch b ON c.branch_id = b.id JOIN commerce co ON c.commerce_id = co.id WHERE b.id = $1 ", id)

	if err != nil {
		return nil, fmt.Errorf("error fetching campaign data: %v", err)
	}
	return result, nil
}

func (b *BranchRepository) CreateBranchCampaign(campaign domain.Campaign) (*domain.Campaign, error) {

	dbConnection, err := b.dbConnection.GetConnection()

	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}

	_, err = dbConnection.NamedExec("INSERT INTO campaign(commerce_id, branch_id, campaign_name, campaign_type, campaign_multiplier, campaign_percent_bonus, start_date, end_date) VALUES(:commerce_id, :branch_id, :campaign_name, :campaign_type, :campaign_multiplier,:campaign_percent_bonus, :start_date, :end_date)", &campaign)

	if err != nil {
		return nil, fmt.Errorf("error saving campaign data: %v", err)
	}

	return &campaign, nil
}
