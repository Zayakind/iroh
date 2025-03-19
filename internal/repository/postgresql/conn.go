package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectToPostgres создает новое соединение с базой данных PostgreSQL
func ConnectToPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to open connection to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}
