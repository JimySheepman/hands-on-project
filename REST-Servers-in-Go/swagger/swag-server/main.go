package main

import (
	"log"
	"net/http"

	"example.com/internal/swagger"
)

func main() {
	log.Printf("Server started")
	router := swagger.NewRouter()
	log.Fatal(http.ListenAndServe("localhost:"+"3000", router))
}
