package services

import (
	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/Edmartt/loyalty-system/internal/core/ports"
)

type CommerceService struct {
	CommerceRepository ports.Commerce
	CampaignRepository ports.Campaign
}

func NewCommerceService(commerceRepository ports.Commerce, campaignRepository ports.Campaign) *CommerceService {
	return &CommerceService{
		CommerceRepository: commerceRepository,
		CampaignRepository: campaignRepository,
	}
}

func (c *CommerceService) CreateCommerce(commerce domain.Commerce) (*domain.Commerce, error) {
	return c.CommerceRepository.SaveCommerce(commerce)
}

func (c *CommerceService) GetCommerceCampaign(campaignID string) ([]domain.CommerceCampaign, error) {
	return c.CampaignRepository.ReadCommerceCampaign(campaignID)
}

func (c *CommerceService) SetCommerceCampaign(campaign domain.Campaign) (*domain.Campaign, error) {
	return c.CampaignRepository.SaveCampaign(campaign)
}

func (c *CommerceService) GetCommercePointsXBuy(id string) (*domain.Commerce, error) {
	return c.CommerceRepository.ReadPointsXBuy(id)
}

func (c *CommerceService) SetCommercePointsXBuy(id string, commercePoint int32) (string, error) {
	return c.CommerceRepository.SavePointsXBuy(id, commercePoint)
}

func (c *CommerceService) SetCommerceValueXPoint(id string, money int64) (string, error) {
	return c.CommerceRepository.SaveValueXPoint(id, money)
}

func (c *CommerceService) GetCommerceValueXPoint(id string) (*domain.Commerce, error) {
	return c.CommerceRepository.ReadValueXPoint(id)
}
