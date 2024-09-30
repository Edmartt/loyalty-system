package services

import (
	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/Edmartt/loyalty-system/internal/core/ports"
)

type BranchService struct {
	campaignRepository ports.Campaign
	branchRepository   ports.Branch
}

func (b *BranchService) CreateBranch(branch domain.Branch) (*domain.Branch, error) {
	return b.branchRepository.CreateBranch(branch)
}

func (b *BranchService) GetBranchCampaign(id string) (*domain.BranchCampaign, error) {
	return b.campaignRepository.ReadBranchCampaign(id)
}

func (b *BranchService) CreateBranchCampaig(campaign domain.Campaign) (*domain.Campaign, error) {
	return b.campaignRepository.SaveCampaign(campaign)
}
