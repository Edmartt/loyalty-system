package services

import (
	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/Edmartt/loyalty-system/internal/core/ports"
)

type CampaignService struct {
	CampaignRepository ports.Campaign
}

func NewCampaignService(campaignRepository ports.Campaign) *CampaignService {
	return &CampaignService{
		CampaignRepository: campaignRepository,
	}
}

func (c *CampaignService) CreateCampaign(campaign domain.Campaign) (*domain.Campaign, error) {
	return c.CampaignRepository.SaveCampaign(campaign)
}
