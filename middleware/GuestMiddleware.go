package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Guest(c *gin.Context) {
	// check is the requester already authenticated
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString != "" {
		result := gin.H{
			"message": "Authorization Header is not null",
		}
		c.JSON(http.StatusNotAcceptable, result)
		c.Abort()
	}
}