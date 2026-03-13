package routes

import (
	"server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func FacilityRoutes(r *gin.Engine) {

	r.POST("/facility", controllers.CreateFacilities)
	r.GET("/facility", controllers.GetFacilities)
	r.GET("/facility/:facility_code", controllers.GetFacilityById)

}
