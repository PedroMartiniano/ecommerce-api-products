package configs

import (
	"database/sql"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	logger = newLogger()
	DB     *sql.DB
	Redis  *redis.Client
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

	Redis = initRedisClient()
}

func GetLogger() *Logger {
	return logger
}
