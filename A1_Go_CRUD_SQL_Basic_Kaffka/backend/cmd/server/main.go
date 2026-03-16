package main

import (
	"server/internal/configs"
	"server/internal/controllers"
	"server/internal/kafka/consumers"
	"server/internal/kafka/producers"
	"server/internal/middlewares"
	"server/internal/repositories"
	"server/internal/routes"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	cfg := configs.Load()

	configs.ConnectDatabase(cfg)
	configs.RunMigrations(cfg)

	facilityRepo := repositories.CreateNewFacilityRepo(configs.DB);
	orderRepo := repositories.CreateNewOrderRepo(configs.DB);

	facilityService := services.CreateNewFacilityService(facilityRepo);
	orderService := services.CreateNewOrderService(orderRepo);

	facilityController := controllers.CreateNewFacilityController(facilityService);
	orderController := controllers.CreateNewOrderController(orderService, facilityService);



	producers.InitializeOrderProducer()
	defer producers.CloseOrder()

	go Consumers.InitializeOrderConsumer()
	r.Use(middlewares.LoggerBro())


	
	routes.FacilityRoutes(r, facilityController)
	routes.OrderRoutes(r, orderController)
	r.Run(":5000")
}
