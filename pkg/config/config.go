package config

import (
	"log"
	"os"
)

func LoadConfig() {
	if err := os.Setenv("DATABASE_URL", "your_db_connection_string"); err != nil {
		log.Fatal(err)
	}
}

func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}
