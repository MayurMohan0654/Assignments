package services

import (
	"server/models"
	"server/repositories"
)


func CreateFacility(facilities *models.Facilities) error{
	err := repositories.CreateFacility(facilities)
	return err
}

func ExistsFacility(code string) bool{
	count := repositories.ExistsFacility(code);
	return count > 0
}