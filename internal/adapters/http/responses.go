package http

import (
	"github.com/Edmartt/loyalty-system/internal/adapters/http/dto"
	"github.com/Edmartt/loyalty-system/internal/core/domain"
)

type HttpResponse struct {
	Response string `json:"response"`
}

type HttpCustomerCreated struct {
	Message  string          `json:"message"`
	Response dto.CustomerDTO `json:"customer"`
}

type HttpCommerceCreated struct {
	Message  string          `json:"message"`
	Response dto.CommerceDTO `json:"commerce"`
}

type HttpCommerceCampaignOk struct {
	Response []domain.Campaign `json:"campaigns"`
}

type HttpCampaignCreated struct {
	Message  string          `json:"message"`
	Response dto.CampaignDTO `json:"campaign"`
}

type HttpBranchCreated struct {
	Message  string        `json:"message"`
	Response dto.BranchDTO `json:"branch"`
}

type HttpBranchCampaignOk struct {
	Response []dto.BranchCampaign `json:"campaigns"`
}

type HttpBranchCampaignCreated struct {
	Message  string             `json:"message"`
	Response dto.BranchCampaign `json:"branch_campaign"`
}
