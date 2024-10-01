package gin

import (
	"net/http"

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
