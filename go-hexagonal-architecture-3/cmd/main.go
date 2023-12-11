package main

import (
	"fmt"
	"log"
	"os"

	"Hexagonal-Architecture/internal/adapters/cache"
	"Hexagonal-Architecture/internal/adapters/handler"
	"Hexagonal-Architecture/internal/adapters/repository"
	"Hexagonal-Architecture/internal/core/domain"
	"Hexagonal-Architecture/internal/core/services"
	"Hexagonal-Architecture/internal/logger"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	msgService     *services.MessengerService
	userService    *services.UserService
	paymentService *services.PaymentService
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	redisCache, err := cache.NewRedisCache("127.0.0.1:6379", "")
	if err != nil {
		panic(err)
	}

	logger.SetupLogger()

	// Create or modify the database tables based on the model structs found in the imported package
	db.AutoMigrate(&domain.Message{}, &domain.User{}, &domain.Payment{})

	store := repository.NewDB(db, redisCache)

	msgService = services.NewMessengerService(store)
	userService = services.NewUserService(store)
	paymentService = services.NewPaymentService(store)

	InitRoutes()
}

func InitRoutes() {
	router := gin.Default()
	router2 := gin.Default()

	pprof.Register(router)
	pprof.Register(router2)

	v1 := router.Group("/v1")

	messageHandler := handler.NewMessageHandler(*msgService)
	v1.GET("/messages/:id", messageHandler.ReadMessage)
	v1.GET("/messages", messageHandler.ReadMessages)
	v1.POST("/messages", messageHandler.CreateMessage)
	v1.PUT("/messages/:id", messageHandler.UpdateMessage)
	v1.DELETE("/messages/:id", messageHandler.DeleteMessage)

	userHandler := handler.NewUserHandler(*userService)
	v1.GET("/users/:id", userHandler.ReadUser)
	v1.GET("/users", userHandler.ReadUsers)
	v1.POST("/users", userHandler.CreateUser)
	v1.PUT("/users", userHandler.UpdateUser)
	v1.DELETE("/users", userHandler.DeleteUser)

	v1.POST("/login", userHandler.LoginUser)
	v1.POST("/membership/webhooks", userHandler.UpdateMembershipStatus)

	v2 := router2.Group("/v2")
	paymentHandler := handler.NewPaymentHandler(*paymentService)
	v2.POST("/create-checkout-session", paymentHandler.CreateCheckoutSession)

	err := router.Run(":4242")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
