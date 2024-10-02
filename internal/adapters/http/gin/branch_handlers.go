package gin

import (
	"net/http"
	"time"

	localHttp "github.com/Edmartt/loyalty-system/internal/adapters/http"
	"github.com/Edmartt/loyalty-system/internal/adapters/http/dto"
	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/Edmartt/loyalty-system/internal/core/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BranchHandler struct {
	BranchService services.BranchService
}

func (b *BranchHandler) PostBranch(context *gin.Context) {

	branchDTO := dto.BranchDTO{}
	response := localHttp.HttpResponse{}
	branchID := uuid.NewString()

	err := context.BindJSON(&branchDTO)

	if err != nil {
		response.Response = "invalid request: " + err.Error()
		context.JSON(http.StatusBadRequest, response)
		return
	}

	if branchDTO.CommerceID == "" || branchDTO.Address == "" {
		response.Response = "missing required information"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	branch := domain.Branch{
		ID:         branchID,
		CommerceID: branchDTO.CommerceID,
		Address:    branchDTO.Address,
	}

	branchResult, err := b.BranchService.CreateBranch(branch)

	if err != nil {
		response.Response = "error creating branch" + err.Error()
		context.JSON(http.StatusBadRequest, response)
		return
	}

	branchDTO.ID = branchResult.ID

	withObjectResponse := localHttp.HttpBranchCreated{
		Message:  "created",
		Response: branchDTO,
	}

	context.JSON(http.StatusCreated, withObjectResponse)

}

func (b *BranchHandler) GetBranchCampaign(context *gin.Context) {

	id := context.Param("id")

	response := localHttp.HttpResponse{}

	if id == "" {
		response.Response = "empty id"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	branchCampaigns, err := b.BranchService.GetBranchCampaign(id)

	if err != nil {
		response.Response = "error fetching data with this id: " + err.Error()
		context.JSON(http.StatusBadRequest, response)
		return
	}

	jsonBranchCampaigns := localHttp.HttpBranchCampaignOk{
		Response: branchCampaigns,
	}

	context.JSON(http.StatusOK, jsonBranchCampaigns)
}

func (b *BranchHandler) CreateBranchCampaign(context *gin.Context) {

	branchCampaignDTO := dto.BranchCampaign{}
	response := localHttp.HttpResponse{}
	campaignID := uuid.NewString()

	err := context.BindJSON(&branchCampaignDTO)

	if err != nil {
		response.Response = "invalid request: " + err.Error()
		context.JSON(http.StatusBadRequest, response)
		return
	}

	if branchCampaignDTO.BranchID == "" || branchCampaignDTO.CampaignName == "" || branchCampaignDTO.StartDate == "" || branchCampaignDTO.EndDate == "" {

		response.Response = "missing required information"
		context.JSON(http.StatusBadRequest, response)
		return
	}
	branchCampaignDTO.CampaignID = campaignID

	campaign := domain.Campaign{
		ID:                   branchCampaignDTO.CampaignID,
		CommerceID:           branchCampaignDTO.CommerceID,
		BranchID:             branchCampaignDTO.BranchID,
		CampaignName:         branchCampaignDTO.CampaignName,
		CampaignType:         branchCampaignDTO.CampaignType,
		CampaignMultiplier:   branchCampaignDTO.CampaignMultiplier,
		CampaignPercentBonus: branchCampaignDTO.CampaignPercentBonus,
	}

	startDateString, err := time.Parse("2006-01-02", branchCampaignDTO.StartDate)

	if err != nil {
		response.Response = "wrong date format"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	campaign.StartDate = startDateString

	endDateString, err := time.Parse("2006-01-02", branchCampaignDTO.EndDate)

	if err != nil {
		response.Response = "wrong date format"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	campaign.EndDate = endDateString

	_, err = b.BranchService.CreateBranchCampaign(campaign)

	if err != nil {
		response.Response = "error creating campaign"
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	//branchCampaignDTO.ID = campaignResult.ID

	withObjectResponse := localHttp.HttpBranchCampaignCreated{
		Message:  "created",
		Response: branchCampaignDTO,
	}

	context.JSON(http.StatusCreated, withObjectResponse)
}
