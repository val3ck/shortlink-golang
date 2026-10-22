package environment

import (
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
)

func EnvGet(key string) string {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error loading .env files.")
	}
	envKey, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal("Error, unavailable .env variable")
	}
	return envKey
}
