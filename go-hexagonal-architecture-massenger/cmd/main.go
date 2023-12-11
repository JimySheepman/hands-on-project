package main

import (
	"Messenger/internal/adapters/handler"
	"Messenger/internal/adapters/repository"
	"Messenger/internal/core/services"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	repo        = flag.String("db", "postgres", "Database for storing messages")
	redisHost   = "localhost:6379"
	httpHandler *handler.HTTPHandler
	svc         *services.MessengerService
)

func main() {
	flag.Parse()

	fmt.Printf("Application running using %s\n", *repo)
	switch *repo {
	case "redis":
		store := repository.NewMessengerRedisRepository(redisHost)
		svc = services.NewMessengerService(store)
	default:
		store := repository.NewMessengerPostgresRepository()
		svc = services.NewMessengerService(store)
	}

	InitRoutes()

}

func InitRoutes() {
	router := gin.Default()
	handler := handler.NewHTTPHandler(*svc)
	router.GET("/messages/:id", handler.ReadMessage)
	router.GET("/messages", handler.ReadMessages)
	router.POST("/messages", handler.SaveMessage)
	router.Run(":5000")
}
