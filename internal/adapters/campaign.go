package adapters

import (
	"fmt"
	"log"

	"github.com/Edmartt/loyalty-system/internal/core/domain"
	ports "github.com/Edmartt/loyalty-system/internal/core/ports/database"
)

type CampaignRepository struct {
	dbConnection ports.IConnection
}

func NewCampaignRepository(dbConnection ports.IConnection) *CampaignRepository {
	return &CampaignRepository{
		dbConnection: dbConnection,
	}
}

func (c *CampaignRepository) SaveCampaign(campaign domain.Campaign) (*domain.Campaign, error) {
	dbConnection, err := c.dbConnection.GetConnection()

	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}

	_, err = dbConnection.NamedExec("INSERT INTO campaign(commerce_id, branch_id, campaign_name, campaign_type, campaign_multiplier, campaign_percent_bonus, start_date, end_date) VALUES(:commerce_id, :branch_id, :campaign_name, :campaign_type, :campaign_multiplier,:campaign_percent_bonus, :start_date, :end_date)", &campaign)

	if err != nil {
		return nil, fmt.Errorf("error saving campaign data: %v", err)
	}

	return &campaign, nil
}

func (c *CampaignRepository) ReadCommerceCampaign(id string) ([]domain.CommerceCampaign, error) {

	dbConnection, err := c.dbConnection.GetConnection()

	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %v", err)
	}

	var result []domain.CommerceCampaign

	err = dbConnection.Select(&result, "SELECT c.id AS campaign_id, c.campaign_name, c.campaign_multiplier, c.campaign_percent_bonus, c.start_date, c.end_date, co.id AS commerce_id, co.name AS commerce_name, co.points_x_buy, co.value_x_point FROM campaign c JOIN commerce co ON c.commerce_id = co.id WHERE co.id = $1", id)

	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error fetching campaign data: %v", err)
	}

	return result, nil
}

func (c *CampaignRepository) ReadBranchCampaign(id string) ([]domain.BranchCampaign, error) {

	dbConnection, err := c.dbConnection.GetConnection()

	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %v", err)
	}

	var result []domain.BranchCampaign

	err = dbConnection.Select(&result, "SELECT c.id AS campaign_id, c.campaign_name, c.campaign_multiplier, c.campaign_percent_bonus, c.start_date, c.end_date, b.id AS branch_id, b.address, co.id AS commerce_id, co.name AS commerce_name, co.points_x_buy, co.value_x_point FROM campaign c JOIN branch b ON c.branch_id = b.id JOIN commerce co ON c.commerce_id = co.id WHERE b.id = $1 ", id)

	if err != nil {
		return nil, fmt.Errorf("error fetching campaign data: %v", err)
	}
	return result, nil
}
