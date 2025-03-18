package repository

import (
	"database/sql"
	"fmt"
	"learn-back/internal/models"

	_ "github.com/lib/pq"
)

// AirportRepository интерфейс для работы с пользователями
type AirportRepository interface {
	GetAllAirports() ([]models.Airport, error)
}

type postgresAirportRepository struct {
	db *sql.DB
}

// NewAirportRepository создает новый экземпляр AirportRepository
func NewAirportRepository(db *sql.DB) *postgresAirportRepository {
	return &postgresAirportRepository{db: db}
}

// GetAllAirports реализует получение всех аэропортав из базы данных
func (repo *postgresAirportRepository) GetAllAirports() ([]models.Airport, error) {
	rows, err := repo.db.Query("SELECT a.airport_code, a.airport_name, a.city, a.latitude, a.longitude, a.timezone FROM bookings.airports a")
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var airports []models.Airport
	for rows.Next() {
		a := models.Airport{}
		if err := rows.Scan(&a.Code, &a.Name, &a.City, &a.Longtitude, &a.Latitude, &a.Timezone); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		airports = append(airports, a)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return airports, nil
}
