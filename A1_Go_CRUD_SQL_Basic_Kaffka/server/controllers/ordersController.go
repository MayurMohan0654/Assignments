package controllers

import (
	"net/http"

	"server/configs"
	"server/models"

	"github.com/gin-gonic/gin"
)

func existsOrder(code string) bool {
	var count int
	response := configs.DB.Raw("SELECT count(*) from orders where id = ?", code).Scan(&count)
	if response.Error != nil {
		return false
	}
	return count > 0
}

func CreateOrders(c *gin.Context) {
	var orders models.Orders

	if err := c.ShouldBindJSON(&orders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if existsOrder(orders.ID) {
		c.JSON(http.StatusConflict, gin.H{"err": "409 Conflict", "msg": orders.ID + " already exists"})
		return
	}

	if !existsFacility(orders.FacilityCode) {
		c.JSON(http.StatusNotFound, gin.H{"err": "404 Not Found", "msg": "No such facility with Facility_code:- " + orders.FacilityCode})
		return
	}

	configs.DB.Create(&orders)

	configs.DB.Raw("select * from facilities where code = ?", orders.FacilityCode).Scan(&orders.Facility)

	c.JSON(http.StatusCreated, orders)

}

// we need to use facility table to fetch the facilities first and then we will find the user in order table which have common facility.code from faciliity table and facilitycode in orders table in common.
//its like join on facility.Code = orders.facilityCode

func GetOrders(c *gin.Context) {
	var orders []models.Orders

	configs.DB.Preload("Facility").Find(&orders)

	c.JSON(http.StatusOK, orders)
}

func GetOrderById(c *gin.Context) {
	var orders []models.Orders

	id := c.Param("order_id")
	result := configs.DB.Raw("Select * from orders where id = ?", id).Scan(&orders)
	
	if result.RowsAffected == 0{
		c.JSON(http.StatusNotFound, gin.H{"err": "404 Not found", "msg" : "No order with id " + id + " was found."})
		return
	}

	c.JSON(http.StatusOK, orders)

}
