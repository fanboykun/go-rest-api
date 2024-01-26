package controllers

import (
	"net/http"

	"rest-api/configs"
	"rest-api/models"
	"rest-api/requests"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

)


func (db_conn *Database) RegisterHandler(c *gin.Context) {
	// bind new user data
	var req requests.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		// error binding request
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Parameter Miss Match",
		})
		return
	}

	db := configs.DBConnect()

	var user models.User
	if db.Where("username = ?", req.Username).First(&user).RowsAffected != 0 {
		// username exists
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Username taken",
		})
		c.Abort()
		return
	}

	user.Username = req.Username
	user.Password = req.Password

	if err := db.Create(&user).Error; err != nil {
		// error saving user
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	// if user created
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"token":  token,
	})
}

func (db_conn *Database) LoginHandler(c *gin.Context) {
	var existingUser models.User
	var req requests.LoginRequest

	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Parameter Miss Match",
		})
	}

	db := configs.DBConnect()
	if db.Where("username = ?", req.Username).First(&existingUser).RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No User Exists With Given Username",
		})
		c.Abort()
		return
	}
	if req.Username != existingUser.Username && req.Password != existingUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Wrong Username or Password",
		})
		c.Abort()
		return
	}

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"token":  token,
	})
}

func (db_conn *Database) LogoutHandler(c *gin.Context) {
	var loggedUser models.User
	var req requests.LogoutRequest

	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Parameter Miss Match",
		})
	}
	db := configs.DBConnect()
	if db.Where("username = ?", req.Username).First(&loggedUser).RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No User Exists With Given Username",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Logged Out",
	})

}
