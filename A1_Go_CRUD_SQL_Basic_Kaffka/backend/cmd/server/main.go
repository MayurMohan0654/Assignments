package main

import (
	"server/configs"
	"server/models"
	"server/routes"

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
