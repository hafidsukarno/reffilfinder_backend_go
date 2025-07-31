package routes

import (
	"github.com/gin-gonic/gin"
	"reffil_finder/controllers"
	"reffil_finder/middleware"
)

func TransaksiRoutes(router *gin.Engine) {
	transaksiGroup := router.Group("/transaksi")
	transaksiGroup.Use(middleware.AuthMiddleware())
	{
		transaksiGroup.POST("/", controllers.CreateTransaksi)
		transaksiGroup.GET("/", controllers.GetTransaksiUser)
		transaksiGroup.PUT("/detail/:id", controllers.UpdateStatusDetail)
		transaksiGroup.DELETE("/:id", controllers.DeleteTransaksi)
	}
}
