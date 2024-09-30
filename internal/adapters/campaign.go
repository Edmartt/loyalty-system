package adapters

import (
	"fmt"

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

	return &campaign, nil
}

func (c *CampaignRepository) ReadCommerceCampaign(id string) (*domain.Campaign, error) {

}

func (c *CampaignRepository) ReadBranchCampaign(id string) (*domain.Branch, error) {

}
