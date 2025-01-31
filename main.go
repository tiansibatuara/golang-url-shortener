package main

import (
	"fmt"
	"url_shortener/config"
	"url_shortener/handlers"
	"url_shortener/models"
	"url_shortener/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	//Conn to Postgres
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	//Auto migrate
	db.AutoMigrate(&models.Url{})

	//Init repo and handlers
	repo := repository.NewURLRepository(db)
	handler := handlers.NewUrlHandler(repo)

	//Router
	r := gin.Default()

	r.POST("/shorten", handler.CreateShortURL)
	r.GET("/shorten/:shortCode", handler.GetOriginalURL)
    r.PUT("/shorten/:shortCode", handler.UpdateShortURL)
    r.DELETE("/shorten/:shortCode", handler.DeleteShortURL)
    r.GET("/shorten/:shortCode/stats", handler.GetStats)
    

	r.Run(":8080")
}