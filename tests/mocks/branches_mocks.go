package mocks

import (
	"fmt"

	"github.com/Edmartt/loyalty-system/internal/adapters/http/dto"
	"github.com/Edmartt/loyalty-system/internal/core/domain"
)

type MockBranchRepository struct {
}

func (m *MockBranchRepository) CreateBranch(branch domain.Branch) (*domain.Branch, error) {

	if branch.ID == "" || branch.CommerceID == "" || branch.Address == "" {
		return nil, fmt.Errorf("missing required information: ")
	}
	return &branch, nil
}

func (m *MockBranchRepository) ReadBranchCampaign(branchID string) ([]dto.BranchCampaign, error) {
	if branchID == "" {
		return nil, fmt.Errorf("ID is required")
	}

	branches := []domain.Branch{
		{
			ID:         "c6f9236a-0805-4c56-aef1-0a717707659e",
			CommerceID: "948d39e1-ed23-4177-bdd2-71df929c3bad",
			Address:    "cl-12 av main street #45",
		},
		{
			ID:         "35816d5d-7648-4a9e-be74-2e463963f180",
			CommerceID: "948d39e1-ed23-4177-bdd2-71df929c3bad",
			Address:    "elm av. #56-67",
		},
		{
			ID:         "5bb6e33c-3673-4488-a00a-735e53d09b85",
			CommerceID: "2278e590-478e-4a33-a01c-36dd20597a4e",
			Address:    "av 26 panamerica #12-54",
		},
	}

	catched := false

	for _, v := range branches {
		if branchID == v.ID {
			catched = true
			break
		}
	}

	if !catched {
		return nil, fmt.Errorf("branchID: %s not found", branchID)
	}

	branchCampaign := []dto.BranchCampaign{
		{
			CampaignID:           "419341b4-2016-40fc-bb74-62b3e0efb66a",
			CampaignName:         "extra points",
			CampaignType:         "points",
			CampaignMultiplier:   0,
			CampaignPercentBonus: 0,
			PointsXBuy:           1,
			ValueXPoint:          600,
			StartDate:            "2024-10-15",
			EndDate:              "2024-10-25",
			BranchID:             "35816d5d-7648-4a9e-be74-2e463963f180",
			CommerceID:           "948d39e1-ed23-4177-bdd2-71df929c3bad",
			CommerceName:         "Texaco",
			Address:              "elm street",
		},
	}

	return branchCampaign, nil
}

func (m *MockBranchRepository) CreateBranchCampaign(campaign domain.Campaign) (*domain.Campaign, error) {
	if campaign.ID == "" || campaign.BranchID == "" || campaign.CampaignType == "" || campaign.StartDate.String() == "" || campaign.EndDate.String() == "" {
		return nil, fmt.Errorf("missing required information")
	}

	return &campaign, nil
}
