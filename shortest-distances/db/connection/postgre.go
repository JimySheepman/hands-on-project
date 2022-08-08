package connection

import (
	"database/sql"
	"log"
	"os"
	"shortest-distances/dto"
)

func Connection(dsnDTO dto.PostgreConnectionDTO) (*sql.DB, error) {
	dns, err := dsnDTO.GenerateDNS()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(os.Getenv("DRIVERNAME"), dns)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected!")
	return db, nil
}
