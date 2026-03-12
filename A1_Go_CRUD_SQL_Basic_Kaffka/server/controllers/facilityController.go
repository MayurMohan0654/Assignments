package controllers

import (
	"net/http"

	"server/configs"
	"server/models"

	"github.com/gin-gonic/gin"
)

func existsFacility(code string) bool {
	var count int
	response := configs.DB.Raw("SELECT count(*) from facilities where code = ?", code).Scan(&count)
	if response.Error != nil {
		return false
	}
	return count > 0
}

func CreateFacilities(c *gin.Context) {
	var facilities models.Facilities

	err := c.ShouldBindJSON(&facilities) // we check if comming json is valid Json or not if okk then populate the facility with comming data (i.e. - body).
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if existsFacility(facilities.Code) {
		c.JSON(http.StatusConflict, gin.H{"msg": facilities.Code + " already exists"})
		return
	}

	configs.DB.Create(&facilities)

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
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "no record found"})
	}

	if result.Error != nil {
		println("somthing went wrong")
		c.JSON(http.StatusNotFound, `{"msg": "notfound or somthing went wrong"}`)
	}
	result.Scan(&facilities)
	c.JSON(http.StatusOK, facilities)

}
