package middleware

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	// if token.Valid && err == nil {
	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}