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
		c.JSON(http.StatusConflict, gin.H{"msg": orders.ID + " already exists"})
		return
	}
	
	configs.DB.Create(&orders)
	c.JSON(http.StatusOK, orders)

}

func GetOrders(c *gin.Context) {
	var orders []models.Orders

	configs.DB.Preload("Facility").Find(&orders) // we need to use facility table to fetch the facilities first and then we will find the user in order table which have common facility.code from faciliity table and facilitycode in orders table in common.
	//its like join on facility.Code = orders.facilityCode

	c.JSON(http.StatusOK, orders)
}

func GetOrdersById(c *gin.Context) {
	var orders []models.Orders

	id := c.Param("order_id")
	configs.DB.Raw("Select * from orders where id = ?", id).Scan(&orders)

	c.JSON(http.StatusOK, orders)
}
