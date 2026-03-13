package repositories

import (
	"server/configs"
	"server/models"
)

func ExistsFacility(code string) int64{
	var count int64;
	configs.DB.Raw("SELECT count(*) from facilities where code = ?", code).Scan(&count);
	return count;
}

func CreateFacility(faciliity *models.Facilities) error{
	result := configs.DB.Create(faciliity)
	return result.Error
}
