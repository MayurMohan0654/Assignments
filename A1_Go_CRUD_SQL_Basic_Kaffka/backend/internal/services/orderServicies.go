package services

import (
	"log"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/kafkaProducers"
)

func CreateOrder(order *models.Orders) error {
	err := repositories.CreateOrder(order)

	if(err != nil){
		return err;
	}

	var savedOrder models.Orders;

	repositories.GetOrderById(&savedOrder, order.ID);
	
	kafkaErr := producers.SendOrderToKAfka(&savedOrder)

	if(kafkaErr != nil){
		log.Printf("order saved in db but faild to send to kafka: %v")
	}

	return nil;
}

func ExistsOrder(code string) bool {
	exist := repositories.ExistsOrder(code)
	return exist
}

func GetAllOrderes(order *[]models.Orders) {
	repositories.GetAllOrderes(order)
}


func GetOrderById(order *models.Orders, id string) int64{
	return repositories.GetOrderById(order, id);
}