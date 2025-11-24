package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	ServerPort string
	DBDSN      string
)

func Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// TODO: Обработка случая с пустыми значениями
	ServerPort = os.Getenv("SERVER_PORT")
	DBDSN = os.Getenv("DB_DSN")

	return nil
}
