package ports

import "github.com/Edmartt/loyalty-system/internal/core/domain"

type Campaign interface {
	SaveCampaign(campaign domain.Campaign) (*domain.Campaign, error)
	ReadCommerceCampaign(id string) (*domain.Campaign, error)
	ReadBranchCampaign(id string) (*domain.Branch, error)
}
