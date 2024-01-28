package main

import (
	// "rest-api/configs"

	"github.com/gin-gonic/gin"

	"rest-api/controllers"
	"rest-api/middleware"
	"rest-api/models"
	// _ "github.com/go-sql-driver/mysql"
)

func main() {
	models.ConnectDataBase()

	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.Use(middleware.JwtAuthMiddleware()).GET("/me", controllers.CurrentUser)
	// router.Use(middleware.JwtAuthMiddleware()).POST("/logout", Database.LogoutHandler)
	
	router.Use(middleware.JwtAuthMiddleware()).GET("/person/:id", controllers.GetPerson)
	// router.Use(middleware.JwtAuthMiddleware()).GET("/persons", controllers.GetPersons)
	// router.Use(middleware.JwtAuthMiddleware()).POST("/person", controllers.CreatePerson)
	// router.Use(middleware.JwtAuthMiddleware()).PUT("/person", controllers.UpdatePerson)
	// router.Use(middleware.JwtAuthMiddleware()).DELETE("/person/:id", controllers.DeletePerson)


	router.Run(":5500")


}

