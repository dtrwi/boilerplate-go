package driver

import (
	"github.com/jmoiron/sqlx"
)

// PostgreOption options for postgresql connection
type PostgreSQLOption struct {
	URL         string
	MaxIdleConn int
	MaxOpenConn int
}

// NewPostgreSQL return a client connection handle to a Postgre server.
func NewPostgreSQL(option PostgreSQLOption) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", option.URL)
	if err != nil {
		return db, err
	}

	db.SetMaxIdleConns(option.MaxIdleConn)
	db.SetMaxOpenConns(option.MaxOpenConn)

	return db, nil
}
