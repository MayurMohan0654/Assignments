package services

import (
	"server/models"
	"server/repositories"
)


func CreateOrder(order *models.Orders) error{
	err := repositories.CreateOrder(order)
	return err;
}

func ExistsOrder(code string) bool{
	exist := repositories.ExistsOrder(code);
	return exist
}