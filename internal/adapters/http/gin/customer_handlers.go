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

// @title		HTTP Client for Loyalty System
// @version		1.0
// @description	Client for handling http requests for loyalty system
// @host		localhost:8080
// @BasePath	/api/v1

// GetUser godoc
// @Summary      Get customers from db using the system service layer
// @Description  Through a get request the customer dni is sent to http client
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        dni   path      string  true  "Customer DNI"
// @Success      200  {object}  domain.Customer
// @Failure	 400 {object}  localHttp.HttpResponse
// @Router       /customers/{dni} [get]
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

// PostUser godoc
// @Summary      Creates new customers
// @Description  Creates customers
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        person body dto.CustomerDTO true "Creates customer"
// @Success      201  {object}  localHttp.HttpCustomerCreated
// @Failure	 400 {object}  localHttp.HttpResponse
// @Router       /customers [post]
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
