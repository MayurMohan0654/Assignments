package repositories

import (
	"server/configs"
	"server/models"
)

func ExistsOrder(code string) bool{
	var count int
	configs.DB.Raw("SELECT count(*) from orders where id = ?", code).Scan(&count)
	return count > 0
}

func CreateOrder(order *models.Orders) error{
	result := configs.DB.Create(order);
	return result.Error
}

func GetAllOrderes(orders *[]models.Orders){
	configs.DB.Preload("Facility").Find(orders);	
}

func GetOrderById(order *[]models.Orders, id string) int64{
	result := configs.DB.Preload("Facility").Where("ID = ?", id).First(order)
	return result.RowsAffected;
}