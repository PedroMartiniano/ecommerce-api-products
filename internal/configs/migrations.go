package configs

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitMigrations() error {
	migrationsPath := GetEnv("MIGRATIONS_PATH")
	dbUrl := GetEnv("DATABASE_URL")

	m, err := migrate.New(
		migrationsPath,
		dbUrl,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
