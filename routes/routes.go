package routes

import (
	"reffil_finder/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// contoh protected route:
	// router.GET("/produk", middlewares.AuthMiddleware(), controllers.GetProduk)
}
