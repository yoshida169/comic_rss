package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const createUsersTable = `
CREATE TABLE IF NOT EXISTS users (
	id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	user_name VARCHAR(100) NOT NULL UNIQUE,
	password_hash VARCHAR(255) NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

// Open connects to MySQL using the given DSN
// (e.g. "user:pass@tcp(127.0.0.1:3306)/dbname?parseTime=true"),
// verifies connectivity, and ensures the schema exists.
func Open(dsn string) (*sql.DB, error) {
	database, err := sql.Open("mysql", dsn)
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
