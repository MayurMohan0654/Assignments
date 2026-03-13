package services

import (
	"server/internal/models"
	"server/internal/repositories"
)

func CreateOrder(order *models.Orders) error {
	err := repositories.CreateOrder(order)
	return err
}

func ExistsOrder(code string) bool {
	exist := repositories.ExistsOrder(code)
	return exist
}

func GetAllOrderes(order *[]models.Orders) {
	repositories.GetAllOrderes(order)
}


func GetOrderById(order *[]models.Orders, id string) int64{
	return repositories.GetOrderById(order, id);
}