package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/JimySheepman/go-rest-api/config/env"
	"github.com/JimySheepman/go-rest-api/internal/router"
)

func init() {
	log.SetPrefix("ERROR: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	_, err := env.LoadEnvironmentConfigure(".env")
	if err != nil {
		log.Fatal("Loading .env file failed")
	}

	srv := &http.Server{
		Handler:      http.HandlerFunc(router.Serve),
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		fmt.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listenAndServe failed: %v", err)
		}
	}()
	fmt.Println("server started")

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped")
}
