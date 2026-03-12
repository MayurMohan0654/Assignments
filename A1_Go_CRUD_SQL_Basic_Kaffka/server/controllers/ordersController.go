package controllers

import (
	"net/http"

	"server/configs"
	"server/models"

	"github.com/gin-gonic/gin"
)

func CreateOrders(c *gin.Context) {

	var user models.Orders

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	configs.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}

func GetOrders(c *gin.Context) {

	var users []models.Orders

	configs.DB.Preload("Facility").Find(&users) // we need to use facility table to fetch the facilities first and then we will find the users in order table which have common facility.code from faciliity table and facilitycode in orders table in common.
	//its like join on facility.Code = orders.facilityCode

	c.JSON(http.StatusOK, users)
}
