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

	// configs.DB.Find(&facilities)
	// or
	configs.DB.Raw("Select * from facilities").Scan(&facilities)

	c.JSON(http.StatusOK, facilities)
}

func GetFacilitiesById(c *gin.Context) {
	var facilities models.Facilities

	code := c.Param("facility_code")

	result := configs.DB.Raw("SELECT * from facilities WHERE code = ?", code)

	if result.Error != nil {
		println("somthing went wrong")
		c.JSON(http.StatusNotFound, `{"msg": "notfound or somthing went wrong"}`)
	}
	result.Scan(&facilities)
	c.JSON(http.StatusOK, facilities)

}
