package controllers

import (
	"net/http"

	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	CreateOrders(c *gin.Context)
	GetOrders(c *gin.Context)
	GetOrderById(c *gin.Context)
}

type orderController struct{
	service services.OrderService
	facilityService services.FacilityService
}


func CreateNewOrderController(service services.OrderService, facilityService services.FacilityService) OrderController{
	controllerVar := orderController{}
	controllerVar.service = service;
	controllerVar.facilityService = facilityService
	return &controllerVar;
}

func (cont *orderController) CreateOrders(c *gin.Context) {
	var orders models.Orders

	if err := c.ShouldBindJSON(&orders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	o_exist := cont.service.ExistsOrder(orders.ID)
	if o_exist {
		c.JSON(http.StatusConflict, gin.H{"err": "409 Conflict", "msg": orders.ID + " already exists"})
		return
	}

	f_exists := cont.facilityService.ExistsFacility(orders.FacilityCode)
	if !f_exists {
		c.JSON(http.StatusNotFound, gin.H{"err": "404 Not Found", "msg": "No such facility with Facility_code:- " + orders.FacilityCode})
		return
	}

	c_err := cont.service.CreateOrder(&orders)
	if c_err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"msg": "error creating the order"})
		return
	}

	cont.facilityService.GetFacilityById(&orders.Facility, orders.FacilityCode) //populates the Facility of created order

	c.JSON(http.StatusCreated, orders)

}

// we need to use facility table to fetch the facilities first and then we will find the user in order table which have common facility.code from faciliity table and facilitycode in orders table in common.
//its like join on facility.Code = orders.facilityCode

func  (cont *orderController)  GetOrders(c *gin.Context) {
	var orders []models.Orders
	cont.service.GetAllOrderes(&orders)
	c.JSON(http.StatusOK, orders)
}

func  (cont *orderController) GetOrderById(c *gin.Context) {
	var order models.Orders

	id := c.Param("order_id")

	count := cont.service.GetOrderById(&order, id)

	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"err": "404 Not found", "msg": "No order with id " + id + " was found."})
		return
	}

	c.JSON(http.StatusOK, order)

}
