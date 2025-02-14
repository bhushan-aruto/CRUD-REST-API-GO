package routes

import (
	"github.com/bhushan-aruto/ice_cream_shop_rest_api/controller"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {

	api := router.Group("/api/flavours")
	{

		api.GET("/", controller.GetAllFlavourController)

		api.GET("/:id", controller.GetFlavourByIdController)

		api.POST("/", controller.AddFlavourController)

		api.PUT("/:id", controller.UpdateFlavourController)

		api.DELETE("/:id", controller.DeleteFlavourController)
	}

}
