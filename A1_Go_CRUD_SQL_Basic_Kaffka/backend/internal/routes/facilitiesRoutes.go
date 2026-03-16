package routes

import (
	"server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func FacilityRoutes(r *gin.Engine, cont controllers.FacilityController) {

	r.POST("/facility",  cont.CreateFacilities)
	r.GET("/facility", cont.GetFacilities)
	r.GET("/facility/:facility_code", cont.GetFacilityById)

}
