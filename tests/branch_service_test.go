package tests

import (
	"testing"
	"time"

	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/Edmartt/loyalty-system/internal/core/services"
	"github.com/Edmartt/loyalty-system/tests/mocks"
	"github.com/google/uuid"
)

func TestCreateBranch(t *testing.T) {
	branchRepo := mocks.MockBranchRepository{}

	branchService := services.NewBranchService(&branchRepo)

	newBranch := domain.Branch{
		ID:         uuid.NewString(),
		CommerceID: "948d39e1-ed23-4177-bdd2-71df929c3bad",
		Address:    "elm street",
	}

	branchCreated, err := branchService.CreateBranch(newBranch)

	if err != nil {
		t.Errorf("error creating branch: expected %v and got %v", branchCreated, err)
	}
}

func TestCreateBranchCampaign(t *testing.T) {
	branchRepo := mocks.MockBranchRepository{}
	branchService := services.NewBranchService(&branchRepo)

	initDate := "2024-11-01"
	finishDate := "2024-12-15"
	startDate, err := time.Parse("2006-01-02", initDate)

	if err != nil {
		t.Errorf("wrong date format: expected %v and got %v", startDate, err)
	}

	endDate, err := time.Parse("2006-01-02", finishDate)

	if err != nil {
		t.Errorf("wrong date format: expected %v and got %v", startDate, err)
	}

	newBranchCampaign := domain.Campaign{
		ID:           uuid.NewString(),
		BranchID:     "35816d5d-7648-4a9e-be74-2e463963f180",
		CampaignType: "cashback",
		StartDate:    startDate,
		EndDate:      endDate,
	}
	result, err := branchService.CreateBranchCampaign(newBranchCampaign)

	if err != nil {
		t.Errorf("error saving branch campaign: expected %v and got %v", result, err)
	}

}

func TestGetBranchCampaign(t *testing.T) {
	branchRepo := mocks.MockBranchRepository{}
	branchService := services.NewBranchService(&branchRepo)
	existantID := "c6f9236a-0805-4c56-aef1-0a717707659e"

	branchCampaign, err := branchService.GetBranchCampaign(existantID)

	if err != nil {
		t.Errorf("error getting result: expected: %v and got %v", branchCampaign, err)
	}
}
