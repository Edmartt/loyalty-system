package services

import (
	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/Edmartt/loyalty-system/internal/core/ports"
)

type CommerceService struct {
	commerceRepository ports.Commerce
	campaignRepository ports.Campaign
}

func (c *CommerceService) CreateCommerce(commerce domain.Commerce) (*domain.Commerce, error) {
	return c.commerceRepository.SaveCommerce(commerce)
}

func (c *CommerceService) GetCommerceCampaign(campaignID string) (*domain.Campaign, error) {
	return c.campaignRepository.ReadCommerceCampaign(campaignID)
}

func (c *CommerceService) SetCommerceCampaign(campaign domain.Campaign) (*domain.Campaign, error) {
	return c.campaignRepository.SaveCampaign(campaign)
}
