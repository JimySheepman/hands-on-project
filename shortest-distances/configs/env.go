package configs

import (
	"errors"

	"github.com/joho/godotenv"
)

func ConfigLoad(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return errors.New("Error loading .env file")
	}
	return nil
}
