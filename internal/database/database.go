package database

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Open opens a database without having to take care of the logic, just calling
// the function should give you the connection with nothing wrong occurs.
func Open() (*sql.DB, error) {
	var db *sql.DB
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
