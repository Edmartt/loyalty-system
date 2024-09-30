package ports

import "github.com/Edmartt/loyalty-system/internal/core/domain"

type Campaign interface {
	SaveCampaign(campaign domain.Campaign) (*domain.Campaign, error)
	ReadCommerceCampaign(id string) (*domain.CommerceCampaign, error)
	ReadBranchCampaign(id string) (*domain.BranchCampaign, error)
}
