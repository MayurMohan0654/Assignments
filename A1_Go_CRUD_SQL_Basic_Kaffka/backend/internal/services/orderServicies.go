package services

import (
	"log"
	"server/internal/kafka/producers"
	"server/internal/models"
	"server/internal/repositories"
)

type OrderService interface {
	ExistsOrder(code string) bool
	CreateOrder(order *models.Orders) error
	GetAllOrderes(order *[]models.Orders)
	GetOrderById(order *models.Orders, id string) int64
}
type orderService struct {
	repo repositories.OrderRepository
}

func CreateNewOrderService(repo repositories.OrderRepository) OrderService {
	serviceVar := orderService{}
	serviceVar.repo = repo
	return &serviceVar

}

func (r *orderService) CreateOrder(order *models.Orders) error {
	err := r.repo.CreateOrder(order)

	if err != nil {
		return err
	}

	var savedOrder models.Orders

	r.repo.GetOrderById(&savedOrder, order.ID)

	kafkaErr := producers.SendOrderToKAfka(&savedOrder)

	if kafkaErr != nil {
		log.Printf("order saved in db but faild to send to kafka: %v")
	}

	return nil
}

func (r *orderService) ExistsOrder(code string) bool {
	exist := r.repo.ExistsOrder(code)
	return exist
}

func (r *orderService) GetAllOrderes(order *[]models.Orders) {
	r.repo.GetAllOrderes(order)
}

func (r *orderService) GetOrderById(order *models.Orders, id string) int64 {
	return r.repo.GetOrderById(order, id)
}
