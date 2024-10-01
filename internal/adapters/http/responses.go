package http

import (
	"github.com/Edmartt/loyalty-system/internal/adapters/http/dto"
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
