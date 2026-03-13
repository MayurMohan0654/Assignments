package controllers

import (
	"net/http"

	"server/configs"
	"server/models"
	"server/services"

	"github.com/gin-gonic/gin"
)

func CreateFacilities(c *gin.Context) {
	var facilities models.Facilities

	err := c.ShouldBindJSON(&facilities) // we check if comming json is valid Json or not if okk then populate the facility with comming data (i.e. - body).
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exist := services.ExistsFacility(facilities.Code)

	if exist {
		c.JSON(http.StatusConflict, gin.H{"err": "409 conflict", "msg": facilities.Code + " already exists"})
		return
	}

	c_err := services.CreateFacility(&facilities)
	if c_err != nil{
		c.JSON(http.StatusNotImplemented, gin.H{"msg":"Error creating the facility"})
		return
	}

	c.JSON(http.StatusCreated, facilities)
}

func GetFacilities(c *gin.Context) {
	var facilities []models.Facilities
	configs.DB.Find(&facilities)
	c.JSON(http.StatusOK, facilities)
}

func GetFacilityById(c *gin.Context) {
	var facilities models.Facilities

	code := c.Param("facility_code")

	result := configs.DB.Raw("SELECT * from facilities WHERE code = ?", code)
	result.Scan(&facilities)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"err": "404 Not Found", "msg": "No facility with code " + code + " found."})
		return
	}

	c.JSON(http.StatusOK, facilities)

}
