package gin

import (
	"log"
	"net/http"
	"time"

	localHttp "github.com/Edmartt/loyalty-system/internal/adapters/http"
	"github.com/Edmartt/loyalty-system/internal/adapters/http/dto"
	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/Edmartt/loyalty-system/internal/core/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CampaignHandlers struct {
	CampaignService services.CampaignService
}

func (c *CampaignHandlers) PostCampaign(context *gin.Context) {
	campaignDTO := dto.CampaignDTO{}
	response := localHttp.HttpResponse{}
	campaignID := uuid.NewString()

	err := context.BindJSON(&campaignDTO)

	if err != nil {
		response.Response = "invalid request: " + err.Error()
		context.JSON(http.StatusBadRequest, response)
		return
	}

	if campaignDTO.BranchID == "" || campaignDTO.CommerceID == "" || campaignDTO.CampaignName == "" || campaignDTO.StartDate == "" || campaignDTO.EndDate == "" {

		response.Response = "missing required information"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	campaign := domain.Campaign{
		ID:                   campaignID,
		CommerceID:           campaignDTO.CommerceID,
		BranchID:             campaignDTO.BranchID,
		CampaignName:         campaignDTO.CampaignName,
		CampaignType:         campaignDTO.CampaignType,
		CampaignMultiplier:   campaignDTO.CampaignMultiplier,
		CampaignPercentBonus: campaignDTO.CampaignPercentBonus,
	}

	startDateString, err := time.Parse("2006-01-02", campaignDTO.StartDate)

	if err != nil {
		log.Println("start date", err.Error())
		response.Response = "wrong date format"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	campaign.StartDate = startDateString

	endDateString, err := time.Parse("2006-01-02", campaignDTO.EndDate)

	if err != nil {
		response.Response = "wrong date format"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	campaign.EndDate = endDateString

	campaignResult, err := c.CampaignService.CreateCampaign(campaign)

	if err != nil {
		response.Response = "error creating campaign"
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	campaignDTO.ID = campaignResult.ID

	withObjectResponse := localHttp.HttpCampaignCreated{
		Message:  "created",
		Response: campaignDTO,
	}

	context.JSON(http.StatusCreated, withObjectResponse)
}
