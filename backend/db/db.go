package db

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const createUsersTable = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	user_name VARCHAR(100) NOT NULL UNIQUE,
	password_hash VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

// Open connects to PostgreSQL using the given DSN
// (e.g. "postgres://user:pass@127.0.0.1:5432/dbname?sslmode=disable"),
// verifies connectivity, and ensures the schema exists.
func Open(dsn string) (*sql.DB, error) {
	database, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	database.SetConnMaxLifetime(3 * time.Minute)

	if err := database.Ping(); err != nil {
		database.Close()
		return nil, err
	}

	if _, err := database.Exec(createUsersTable); err != nil {
		database.Close()
		return nil, err
	}

	return database, nil
}
