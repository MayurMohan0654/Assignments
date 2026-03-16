package services

import (
	"server/internal/models"
	"server/internal/repositories"
)

type FacilityService interface {
	CreateFacility(facilities *models.Facilities) error
	ExistsFacility(code string) bool
	GetAllFacilities(here *[]models.Facilities)
	GetFacilityById(facility *models.Facilities, code string) int64
}

type faciliityService struct {
	repo repositories.FaciliityRepository
}

func CreateNewFacilityService(repo repositories.FaciliityRepository) FacilityService {
	servviceVar := faciliityService{}
	servviceVar.repo = repo
	return &servviceVar
}

func (s *faciliityService) CreateFacility(facilities *models.Facilities) error {
	err := s.repo.CreateFacility(facilities)
	return err
}

func (s *faciliityService) ExistsFacility(code string) bool {
	count := s.repo.ExistsFacility(code)
	return count > 0
}

func (s *faciliityService) GetAllFacilities(here *[]models.Facilities) {
	s.repo.GetAllFacilities(here)
}

func (s *faciliityService) GetFacilityById(facility *models.Facilities, code string) int64 {
	count := s.repo.GetFacilityById(facility, code)
	return count
}
