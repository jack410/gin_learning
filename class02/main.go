package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//Get /user
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Get all users"})
	})

	//POST /users
	r.POST("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Create user"})
	})

	//PUT /users
	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"data": "Update user " + id})
	})

	//DELETE /users
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"data": "Delete user " + id})
	})

	r.Run(":8888")
}
