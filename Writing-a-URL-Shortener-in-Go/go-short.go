package main

import (
	"log"
	"net/http"

	"go-short/web"
)

func main() {
	mux := http.NewServeMux()
	web.RegisterHandlers(mux)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("unable to start server: %v\n", err)
	}
}
