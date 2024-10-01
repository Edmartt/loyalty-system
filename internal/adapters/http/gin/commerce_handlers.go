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

type CommerceHandlers struct {
	CommerceService services.CommerceService
}

func (c CommerceHandlers) PostCommerce(context *gin.Context) {

	commerceDTO := dto.CommerceDTO{}
	response := localHttp.HttpResponse{}
	commerceID := uuid.NewString()

	err := context.BindJSON(&commerceDTO)

	if err != nil {
		response.Response = "invalid request: " + err.Error()
		context.JSON(http.StatusBadRequest, response)
		return
	}

	if commerceDTO.Name == "" {
		response.Response = "missing required information"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	commerce := domain.Commerce{
		ID:   commerceID,
		Name: commerceDTO.Name,
	}

	commerceResult, err := c.CommerceService.CreateCommerce(commerce)

	if err != nil {
		response.Response = "error creating commerce: " + err.Error()
		context.JSON(http.StatusBadRequest, response)
		return
	}

	commerceDTO.ID = commerceResult.ID

	withObjectResponse := localHttp.HttpCommerceCreated{
		Message:  "commerce created successfully",
		Response: commerceDTO,
	}

	context.JSON(http.StatusCreated, withObjectResponse)
}

func (c CommerceHandlers) GetCommerceCampaign(context *gin.Context) {
	id := context.Param("id")

	response := localHttp.HttpResponse{}

	if id == "" {
		response.Response = "empty id"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	commerceCampaigns, err := c.CommerceService.GetCommerceCampaign(id)

	if err != nil {
		response.Response = "error fetching data with this id"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	jsonCommerceCampaigns := localHttp.HttpCommerceCampaignOk{
		Response: commerceCampaigns,
	}

	context.JSON(http.StatusOK, jsonCommerceCampaigns)
}

func (c CommerceHandlers) PostCommerceCampaign(context *gin.Context) {

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

	campaignResult, err := c.CommerceService.SetCommerceCampaign(campaign)

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
