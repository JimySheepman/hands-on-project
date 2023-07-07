package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connDB := fmt.Sprintf("postgres://merlins:root@localhost/er?sslmode=disable")

	db, err := sql.Open("postgres", connDB)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	return db, nil
}
