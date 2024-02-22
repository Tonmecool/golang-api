package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURL() string{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env")
	}

	return os.Getenv("MONGOURL")
}

