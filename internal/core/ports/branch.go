package ports

import (
	"github.com/Edmartt/loyalty-system/internal/adapters/http/dto"
	"github.com/Edmartt/loyalty-system/internal/core/domain"
)

type Branch interface {
	CreateBranch(domain.Branch) (*domain.Branch, error)
	CreateBranchCampaign(domain.Campaign) (*domain.Campaign, error)
	ReadBranchCampaign(string) ([]dto.BranchCampaign, error)
}
