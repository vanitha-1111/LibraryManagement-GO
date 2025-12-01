package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewDB connection creates and returns a PostgreSQL database connection using sqlx
func NewDBConnection(conn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
