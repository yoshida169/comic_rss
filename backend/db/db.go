package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

const createUsersTable = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	user_name VARCHAR(100) NOT NULL UNIQUE,
	password_hash VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

// DSN returns the PostgreSQL DSN. DB_DSN takes precedence; otherwise it is built
// from USERNAME / USERPASS / DATABASE in the repository-root .env file
// (the same variables docker-compose.yaml uses). DB_HOST overrides the
// default host:port of 127.0.0.1:5432.
func DSN() (string, error) {
	if dsn := os.Getenv("DB_DSN"); dsn != "" {
		return dsn, nil
	}

	// Read into a map instead of Load so that shell variables such as zsh's
	// USERNAME cannot shadow the values in .env.
	env, err := godotenv.Read("../.env")
	if err != nil {
		return "", fmt.Errorf("DB_DSN is not set and .env could not be read: %w", err)
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.1:5432"
	}
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(env["USERNAME"], env["USERPASS"]),
		Host:     host,
		Path:     "/" + env["DATABASE"],
		RawQuery: "sslmode=disable",
	}
	return u.String(), nil
}

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
