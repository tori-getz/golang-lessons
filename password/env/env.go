package env

import (
	"os"
	"password/app/log"

	"github.com/joho/godotenv"
)

const Encryptkey = "ENCRYPT_KEY"

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Error("Unable to load .env!")
		return
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
