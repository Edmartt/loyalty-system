package gin

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type HTTPServer struct {
	CustomerHandler CustomerHandlers
	CommerceHandler CommerceHandlers
	CampaignHandler CampaignHandlers
}

func (h HTTPServer) setRoutes(router *gin.RouterGroup) {
	router.GET("/customers/:dni", h.CustomerHandler.GetUser)
	router.POST("/customers", h.CustomerHandler.PostUser)

	router.POST("/commerces", h.CommerceHandler.PostCommerce) //new commerce
	router.GET("/commerces/:id/campaigns", h.CommerceHandler.GetCommerceCampaign)
	router.POST("/commerces/:commerce_id/campaigns", h.CampaignHandler.PostCampaign) // campaigns related to whole commerce
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	/*	router.POST("/commerces/:id/campaigns")

		router.POST("/branches") //new branch
		router.GET("/branches/:id/campaigns")
		router.POST("/branches/:id/campaigns")

		router.GET("/commerces/:id") // for points x buy
		router.POST("/commerces/")   // for updating points x buy

		router.GET("/commerces/:id") // for value x point
		router.POST("/commerces/")   // for updating value x point*/

}

func (h HTTPServer) setRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")
	h.setRoutes(apiGroup)

	return router
}

// RunServer starts http server
func (h HTTPServer) RunServer(port string) error {
	router := h.setRouter()
	log.Fatal(router.Run(":" + port))

	return nil
}
