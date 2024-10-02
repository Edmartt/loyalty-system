package services

import (
	"github.com/Edmartt/loyalty-system/internal/adapters/http/dto"
	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/Edmartt/loyalty-system/internal/core/ports"
)

type BranchService struct {
	branchRepository ports.Branch
}

func NewBranchService(branch ports.Branch) *BranchService {
	return &BranchService{
		branchRepository: branch,
	}
}

func (b *BranchService) CreateBranch(branch domain.Branch) (*domain.Branch, error) {
	return b.branchRepository.CreateBranch(branch)
}

func (b *BranchService) GetBranchCampaign(id string) ([]dto.BranchCampaign, error) {
	return b.branchRepository.ReadBranchCampaign(id)
}

func (b *BranchService) CreateBranchCampaign(campaign domain.Campaign) (*domain.Campaign, error) {
	return b.branchRepository.CreateBranchCampaign(campaign)
}
