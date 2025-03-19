package service

import (
	"learn-back/internal/models"
	repository "learn-back/internal/repository/postgresql"
)

type AircraftService struct {
	repo repository.AircraftRepository
}

func NewAircraftService(repo repository.AircraftRepository) *AircraftService {
	return &AircraftService{repo: repo}
}

func (as *AircraftService) GetAllAircrafts() ([]models.Aircraft, error) {
	return as.repo.GetAllAircrafts()
}
