package http

import "github.com/Edmartt/loyalty-system/internal/core/domain"

type HttpResponse struct {
	Response string `json:"response"`
}

type HttpCustomerCreated struct {
	Message  string          `json:"message"`
	Response domain.Customer `json:"customer"`
}
