package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvPort() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("PORT")
}

func EnvMongoDB() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGO_URI")
}
