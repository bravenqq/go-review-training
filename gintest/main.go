// Package main provides ...
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	v1 := r.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.String(200, "login")
		})

	}

	r.Run()
}
