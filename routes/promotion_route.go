package routes

import (
    "reffil_finder/controllers"
	"reffil_finder/middleware"
	"github.com/gin-gonic/gin"
)

func PromotionRoutes(r *gin.Engine) {
    
	promotion := r.Group("/promotions")
	promotion.Use(middleware.AuthMiddleware())
	{
    	promotion.GET("/", controllers.GetPromotionsBySeller)
    	promotion.POST("/", controllers.CreatePromotion)
    	promotion.PUT("/:id", controllers.UpdatePromotion)
    	promotion.DELETE("/:id", controllers.DeletePromotion)
	}

}
