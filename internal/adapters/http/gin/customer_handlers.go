package gin

import (
	"net/http"

	localHttp "github.com/Edmartt/loyalty-system/internal/adapters/http"
	"github.com/Edmartt/loyalty-system/internal/adapters/http/dto"
	"github.com/Edmartt/loyalty-system/internal/core/domain"
	"github.com/Edmartt/loyalty-system/internal/core/services"
	"github.com/gin-gonic/gin"
)

type CustomerHandlers struct {
	CustomerService services.CustomerService
}

func (c CustomerHandlers) GetUser(context *gin.Context) {

	dni := context.Param("dni")

	response := localHttp.HttpResponse{}

	if dni == "" {
		response.Response = "dni empty"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	customer, err := c.CustomerService.GetUser(dni)

	if err != nil {
		response.Response = "wrong dni"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	context.JSON(http.StatusOK, customer)
}

func (c CustomerHandlers) PostUser(context *gin.Context) {

	customerDTO := dto.CustomerDTO{}
	response := localHttp.HttpResponse{}

	err := context.BindJSON(&customerDTO)

	if err != nil {
		response.Response = "invalid request: " + err.Error()
		context.JSON(http.StatusBadRequest, response)
		return
	}

	if customerDTO.UserDNI == "" || customerDTO.Name == "" || customerDTO.LastName == "" || customerDTO.Phone == "" || customerDTO.Email == "" {

		response.Response = "missing required data"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	customer := domain.Customer{
		UserDNI:  customerDTO.UserDNI,
		Name:     customerDTO.Name,
		LastName: customerDTO.LastName,
		Phone:    customerDTO.Phone,
		Email:    customerDTO.Email,
	}

	_, err = c.CustomerService.CreateUser(customer)

	if err != nil {
		response.Response = "error creating user: " + err.Error()
		context.JSON(http.StatusBadRequest, response)
		return
	}

	withObjectResponse := localHttp.HttpCustomerCreated{
		Message:  "customer created successfully",
		Response: customerDTO,
	}
	context.JSON(http.StatusOK, withObjectResponse)
}
