package main

import (
	"gin-mysql-api/config"
	"gin-mysql-api/controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConncetion()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run()
}
