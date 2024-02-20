package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var EnvMongoUser string
var EnvMongoPassword string
var EnvMongoHost string
var EnvMongoPort string

func LoadEnvironments() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	EnvMongoUser = os.Getenv("DB_MONGO_USER")
	EnvMongoPassword = os.Getenv("DB_MONGO_PASS")
	EnvMongoHost = os.Getenv("DB_MONGO_HOST")
	EnvMongoPort = os.Getenv("DB_MONGO_PORT")
}
