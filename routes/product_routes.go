package routes

import (
	"reffil_finder/controllers"
	"reffil_finder/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	product := r.Group("/products")
	product.Use(middleware.AuthMiddleware())
	{
		product.GET("/", controllers.GetAllProducts)
		product.POST("/", controllers.CreateProduct)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
	}
}
