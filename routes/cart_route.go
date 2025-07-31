package routes

import (
	"reffil_finder/controllers"
	"reffil_finder/middleware"
	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine) {
	cart := r.Group("/cart")
	cart.Use(middleware.AuthMiddleware()) 
	{
		cart.GET("/", controllers.GetCart)
		cart.POST("/", controllers.AddToCart)
		cart.PUT("/:id", controllers.UpdateCartItem)
		cart.DELETE("/:id", controllers.DeleteCartItem)
	}

}
