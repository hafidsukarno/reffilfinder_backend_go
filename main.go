package main

import (
	"reffil_finder/config"
	"reffil_finder/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	config.InitDB()
	routes.SetupRoutes(r)
	
	routes.ProductRoutes(r)
	routes.PromotionRoutes(r)
	routes.CartRoutes(r)
	routes.TransaksiRoutes(r)
	r.Run(":8080")
}
