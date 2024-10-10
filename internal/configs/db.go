package configs

import (
	_ "github.com/lib/pq"
	"database/sql"
)

func initPostgres() (*sql.DB, error) {
	conn := GetEnv("DATABASE_URL")
	DB, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	return DB, nil
}
