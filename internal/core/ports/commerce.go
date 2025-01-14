package ports

import "github.com/Edmartt/loyalty-system/internal/core/domain"

type Commerce interface {
	SaveCommerce(domain.Commerce) (*domain.Commerce, error)
	ReadCommerceCampaign(string) ([]domain.Campaign, error)
	CreateCommerceCampaign(domain.Campaign) (*domain.Campaign, error)
	ReadPointsXBuy(string) (*domain.Commerce, error)
	SavePointsXBuy(string, int32) (string, error)
	ReadValueXPoint(id string) (*domain.Commerce, error)
	SaveValueXPoint(string, int64) (string, error)
}
