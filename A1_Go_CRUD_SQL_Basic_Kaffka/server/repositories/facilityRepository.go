package repositories

import (
	"server/configs"
	"server/models"
)

func ExistsFacility(code string) int64 {
	var count int64
	configs.DB.Raw("SELECT count(*) from facilities where code = ?", code).Scan(&count)
	return count
}

func CreateFacility(faciliity *models.Facilities) error {
	result := configs.DB.Create(faciliity)
	return result.Error
}

func GetAllFacilities(here *[]models.Facilities) error {
	result := configs.DB.Find(&here)
	return result.Error
}

func GetFacilityById(facility *models.Facilities, code string) (int64){
	result := configs.DB.Where("code = ?", code).First(facility);
	return result.RowsAffected
}