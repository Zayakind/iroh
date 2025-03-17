package service

import (
	"errors"

	"learn-back/internal/models"
	repository "learn-back/internal/repository/postgresql"
)

// AirportService структура сервиса для работы с аэропортами
type AirportService struct {
	repo repository.AirportRepository
}

// NewAirportrService создает новый экземпляр AirportService
func NewAirportrService(repo repository.AirportRepository) *AirportService {
	return &AirportService{repo: repo}
}

// GetAllAirports возвращает список всех Аэропортов
func (as *AirportService) GetAllAirports() ([]*models.Airport, error) {
	airports, err := as.repo.GetAllAirports()
	if err != nil {
		return nil, errors.New("failed to get all users")
	}
	return airports, nil
}
