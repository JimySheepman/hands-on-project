package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func loginAdmin(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	if username == "admin" && password == "123" {
		cookie := &http.Cookie{
			Name:    "userID",
			Value:   "user_id",
			Expires: time.Now().Add(48 * time.Hour),
		}

		c.SetCookie(cookie)
		return c.String(http.StatusOK, "Login success")
	}

	return c.String(http.StatusUnauthorized, "wrong username or password")
}

func userHandler(c echo.Context) error {
	dataType := c.Param("data")

	// query param
	// /user?username=<username>&name=<name>&surname=<surname>
	username := c.QueryParam("username")
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Username: %s, Name: %s, Surname: %s", username, name, surname))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"username": username,
			"name":     name,
			"surname":  surname,
		})
	}

	return c.String(http.StatusBadRequest, "Only accapted json or string data type")
}

func addUser(c echo.Context) error {
	user := User{}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Success")
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Admin endpoint")
}

func setHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Request().Header.Get("Accept"))
		return next(c)
	}
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("userId")
		if err != nil {
			return err
		}

		if cookie.Value == "user_id" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "wrong cookie")
	}
}

func main() {
	fmt.Println(os.Getenv("GITHUB_USERNAME"))
	e := echo.New()

	// * use all endpoint
	//e.Use(middleware.Logger())
	// * logger with congif
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "method=${method}, uri=${uri}, status=${status}\n",}))
	e.GET("/main", mainHandler)
	// e.Use(setHeader)

	// localhost:8080/admin
	// * use middleware spesificly group middleware
	adminGroup := e.Group("/admin", middleware.Logger())
	// if use group middleware, create group after the use middleware. if it any endpoint before use, it cannot use middleware.
	/* 	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "123" {
			return true, nil
		}
		return false, nil
	})) */

	// localhost:8080/admin/main
	adminGroup.GET("/main", mainAdmin, checkCookie)
	adminGroup.GET("/login", loginAdmin)

	e.GET("/user/:data", userHandler)
	e.POST("/user", addUser)

	e.Logger.Fatal(e.Start(":8080"))
}
