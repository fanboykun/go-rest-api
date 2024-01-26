package main

import (
	"rest-api/configs"
	"rest-api/controllers"
	"rest-api/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := configs.DBConnect()
	Database := &controllers.Database{DB: db}

	router := gin.Default()

	router.POST("/register", middleware.Guest, Database.RegisterHandler)
	router.POST("/login", middleware.Guest, Database.LoginHandler)
	router.POST("/logout", middleware.Auth, Database.LogoutHandler)
	

	router.GET("/person/:id", middleware.Auth, Database.GetPerson)
	router.GET("/persons", middleware.Auth, Database.GetPersons)
	router.POST("/person", middleware.Auth, Database.CreatePerson)
	router.PUT("/person", middleware.Auth, Database.UpdatePerson)
	router.DELETE("/person/:id", middleware.Auth, Database.DeletePerson)

	router.Run(":3000")
}

