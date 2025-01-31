package main

import (
	"github.com/tiansibatuara/golang-url-shortener/initializers"
	"github.com/tiansibatuara/golang-url-shortener/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Url{})
}