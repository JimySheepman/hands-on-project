package dto

import (
	"fmt"
	"os"
	"strconv"
)

type PostgreConnectionDTO struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"Database_name"`
}

func (d *PostgreConnectionDTO) New() PostgreConnectionDTO {
	return PostgreConnectionDTO{
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		User:         os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
	}
}

func (d *PostgreConnectionDTO) GenerateDNS() (string, error) {
	port, err := strconv.Atoi(d.Port)
	if err != nil {
		return "", err
	}

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", d.Host, port, d.User, d.Password, d.DatabaseName)

	return conn, err
}
