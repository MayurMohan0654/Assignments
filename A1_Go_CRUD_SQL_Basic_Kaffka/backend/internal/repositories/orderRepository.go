package repositories

import (
	"server/internal/configs"
	"server/internal/models"

	"gorm.io/gorm"
)

type OrderRepository interface{
	CreateOrder(order *models.Orders) error 
	GetAllOrderes(orders *[]models.Orders) 
	GetOrderById(order *models.Orders, id string) int64
	ExistsOrder(code string) bool

}

type orderRepository struct{
	db *gorm.DB;
}

func CreateNewOrderRepo(db *gorm.DB) OrderRepository{
	repoVar := orderRepository{};
	repoVar.db = db
	return &repoVar
}

func (r *orderRepository) ExistsOrder(code string) bool {
	var count int
	configs.DB.Raw("SELECT count(*) from orders where id = ?", code).Scan(&count)
	return count > 0
}

func (r *orderRepository) CreateOrder(order *models.Orders) error {
	result := configs.DB.Create(order)

	return result.Error
}

func (r *orderRepository) GetAllOrderes(orders *[]models.Orders) {
	configs.DB.Preload("Facility").Find(orders)
}

func (r *orderRepository) GetOrderById(order *models.Orders, id string) int64 {
	result := configs.DB.Preload("Facility").Where("ID = ?", id).First(order)
	return result.RowsAffected
}
