package repository

import (
	"database/sql"
	"fmt"
	"learn-back/internal/models"
)

// type AircraftRepository interface {
//     Create(a *Aircraft) error
//     Get(id int64) (*Aircraft, error)
//     Update(a *Aircraft) error
//     Delete(id int64) error
//     List() ([]*Aircraft, error)
// }

type AircraftRepository interface {
	GetAllAircrafts() ([]*models.Aircraft, error)
}

type postgresAircraftRepository struct {
	db *sql.DB
}

// NewAicraftRepository создает новый экземпляр AirportRepository
func NewAicraftRepository(db *sql.DB) AircraftRepository {
	return &postgresAircraftRepository{db: db}
}

func (repo *postgresAircraftRepository) GetAllAircrafts() ([]*models.Aircraft, error) {
	rows, err := repo.db.Query("SELECT a.aircraft_code, a.model, a.range FROM bookings.aircrafts a")
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var aircrafts []*models.Aircraft
	for rows.Next() {
		a := new(models.Aircraft)
		if err := rows.Scan(&a.Code, &a.Model, &a.Range); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		aircrafts = append(aircrafts, a)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return aircrafts, nil
}
