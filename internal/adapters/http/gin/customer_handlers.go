package gin

import (
	"net/http"

	localHttp "github.com/Edmartt/loyalty-system/internal/adapters/http"
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
