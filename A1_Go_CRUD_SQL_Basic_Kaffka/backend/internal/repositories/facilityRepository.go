package repositories

import (
	"server/internal/models"
	"gorm.io/gorm"
)

type FaciliityRepository interface {
	ExistsFacility(code string) int64
	CreateFacility(facility *models.Facilities) error
	GetAllFacilities(here *[]models.Facilities)
	GetFacilityById(faciliity *models.Facilities, code string) int64
}

type faciliityRepository struct {
	db *gorm.DB
}

func CreateNewFacilityRepo(db *gorm.DB) FaciliityRepository {
	repoVar := faciliityRepository{};
	repoVar.db = db
	return &repoVar
}

func (r *faciliityRepository)  ExistsFacility(code string) int64 {
	var count int64
	r.db.Raw("SELECT count(*) from facilities where code = ?", code).Scan(&count)
	return count
}

func (r *faciliityRepository) CreateFacility(faciliity *models.Facilities) error {
	result := r.db.Create(faciliity)
	return result.Error
}

func (r *faciliityRepository)  GetAllFacilities(here *[]models.Facilities){
	r.db.Find(&here)
}

func (r *faciliityRepository)  GetFacilityById(facility *models.Facilities, code string) int64 {
	result := r.db.Where("code = ?", code).First(facility)
	return result.RowsAffected
}
