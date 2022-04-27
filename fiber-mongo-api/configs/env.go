package configs

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

func EnvMongoURI() string {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGOURI")
}
