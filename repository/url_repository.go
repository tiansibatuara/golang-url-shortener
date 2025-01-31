package repository

import (
	"github.com/tiansibatuara/golang-url-shortener/models"
)

type UrlRepository interface {
	Create(url *models.Url) error
	FindByShortCode(code string) (*models.Url, error)
	Update(url *models.Url) error
	Delete(code string) error
	IncrementAccessCount(code string) error
}
