package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConnection struct {
	File string
}

func DbConfig() DbConnection {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	dbConfig := DbConnection{
		File: os.Getenv("DB_FILE"),
	}

	return dbConfig
}
