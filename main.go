package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	r := gin.Default()

	r.GET("/example", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "example",
		})
	})

	r.Run(":8080")
}