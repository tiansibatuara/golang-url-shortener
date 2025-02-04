package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tiansibatuara/golang-url-shortener/handlers"
	"github.com/tiansibatuara/golang-url-shortener/initializers"
	"github.com/tiansibatuara/golang-url-shortener/repository"
	"github.com/tiansibatuara/golang-url-shortener/router"
	"github.com/tiansibatuara/golang-url-shortener/service"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	db := initializers.DB
	urlRepo := repository.NewUrlRepository(db)
	urlService := service.NewUrlServiceImpl(urlRepo)
	urlHandler := handlers.NewUrlHandler(urlService)

	r := gin.Default()

	router.NewUrlRouter(r, urlHandler)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start server
	r.Run()
}
