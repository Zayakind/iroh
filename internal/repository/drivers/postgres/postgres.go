package postgres

import (
	"database/sql"
	"fmt"

	"learn-back/internal/config"

	_ "github.com/lib/pq"
)

// ConnectToPostgres создает новое соединение с базой данных PostgreSQL
func ConnectToPostgres(cfg *config.Config) (*sql.DB, error) {
	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to open connection to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}
