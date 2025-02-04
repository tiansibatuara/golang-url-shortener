package service

import "github.com/tiansibatuara/golang-url-shortener/models"

type UrlService interface {
	CreateUrl(originalUrl string) (*models.Url, error)
	GetOriginalUrl(code string) (*models.Url, error)
	UpdateUrl(code, newUrl string) (*models.Url, error)
	DeleteUrl(code string) error
	GetStats(code string) (*models.Url, error)
}
