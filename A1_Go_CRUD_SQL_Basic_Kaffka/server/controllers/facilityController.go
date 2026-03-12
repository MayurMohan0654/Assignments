package controllers

import (
	// "fmt"
	"fmt"
	"net/http"

	"server/configs"
	"server/models"

	"github.com/gin-gonic/gin"
)

func CreateFacilities(c *gin.Context) {

	var facilities models.Facilities

	err := c.ShouldBindJSON(&facilities) // we check if comming json is valid Json or not
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := configs.DB.Create(&facilities)
	if response != nil {
		fmt.Println("hello to all", response)
	}

	c.JSON(http.StatusOK, facilities)
}

func GetFacilities(c *gin.Context) {

	var facilities []models.Facilities

	configs.DB.Find(&facilities)

	c.JSON(http.StatusOK, facilities)
}
