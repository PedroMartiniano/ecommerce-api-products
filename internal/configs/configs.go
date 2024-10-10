package configs

import (
	"database/sql"
	"github.com/joho/godotenv"
)

var (
	logger = newLogger()
	DB     *sql.DB
)

func Init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		logger.Errorf("Error loading .env file, error: %v", err)
		panic(err)
	}

	DB, err = initPostgres()
	if err != nil {
		logger.Errorf("Error initializing database, error: %v", err)
		panic(err)
	}

	err = InitMigrations()
	if err != nil {
		logger.Errorf("Error running migrations, error: %v", err)
		panic(err)
	}
}

func GetLogger() *Logger {
	return logger
}
