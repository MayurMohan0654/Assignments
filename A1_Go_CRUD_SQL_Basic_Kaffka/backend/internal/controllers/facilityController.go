package controllers

import (
	"net/http"

	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type FacilityController interface {
	CreateFacilities(c *gin.Context)
	GetFacilities(c *gin.Context)
	GetFacilityById(c *gin.Context)
}

type facilityController struct {
	service services.FacilityService
}

func CreateNewFacilityController(service services.FacilityService) FacilityController {
	controllerVar := facilityController{}
	controllerVar.service = service
	return &controllerVar
}

func (cont *facilityController) CreateFacilities(c *gin.Context) {
	var facilities models.Facilities

	err := c.ShouldBindJSON(&facilities) // we check if comming json is valid Json or not if okk then populate the facility with comming data (i.e. - body).
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exist := cont.service.ExistsFacility(facilities.Code)

	if exist {
		c.JSON(http.StatusConflict, gin.H{"err": "409 conflict", "msg": facilities.Code + " already exists"})
		return
	}

	c_err := cont.service.CreateFacility(&facilities)
	if c_err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"msg": "Error creating the facility"})
		return
	}

	c.JSON(http.StatusCreated, facilities)
}

func (cont *facilityController) GetFacilities(c *gin.Context) {
	var facilities []models.Facilities
	cont.service.GetAllFacilities(&facilities)
	c.JSON(http.StatusOK, facilities)
}

func (cont *facilityController) GetFacilityById(c *gin.Context) {
	var facilities models.Facilities

	code := c.Param("facility_code")

	count := cont.service.GetFacilityById(&facilities, code) // find from database and populate the facility variable (thus sending the pointer)

	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"err": "404 Not Found", "msg": "No facility with code " + code + " found."})
		return
	}

	c.JSON(http.StatusFound, facilities)
}
