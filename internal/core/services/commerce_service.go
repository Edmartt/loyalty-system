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

func (c *CommerceService) GetCommerceCampaign(campaignID string) (*domain.CommerceCampaign, error) {
	return c.campaignRepository.ReadCommerceCampaign(campaignID)
}

func (c *CommerceService) SetCommerceCampaign(campaign domain.Campaign) (*domain.Campaign, error) {
	return c.campaignRepository.SaveCampaign(campaign)
}

func (c *CommerceService) GetCommercePointsXBuy(id string) (*domain.Commerce, error) {
	return c.commerceRepository.ReadPointsXBuy(id)
}

func (c *CommerceService) SetCommercePointsXBuy(id string, commercePoint int32) (string, error) {
	return c.commerceRepository.SavePointsXBuy(id, commercePoint)
}

func (c *CommerceService) SetCommerceValueXPoint(id string, money int64) (string, error) {
	return c.commerceRepository.SaveValueXPoint(id, money)
}

func (c *CommerceService) GetCommerceValueXPoint(id string) (*domain.Commerce, error) {
	return c.commerceRepository.ReadValueXPoint(id)
}
