package main

import (
	"server/internal/configs"
	"server/internal/models"
	"server/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	configs.ConnectDatabase()

	configs.DB.AutoMigrate(&models.Facilities{})
	configs.DB.AutoMigrate(&models.Orders{})

	routes.FacilityRoutes(r)
	routes.OrderRoutes(r)

	r.Run(":5000")
}
