package driver

import (
	"database/sql"
	"os"

	"github.com/lib/pq"
)

// ConnectDB gets SQL_URL-environmental variable and connects to psql
func ConnectDB() (*sql.DB, error) {
	pgURL, err := pq.ParseURL(os.Getenv("SQL_URL"))
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
