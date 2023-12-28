package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dashboard"})
	})

	auth := r.Group("/auth")
	{
		auth.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "auth"})
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
	})
	return r
}
