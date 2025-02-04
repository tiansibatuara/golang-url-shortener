package service

import (
	"errors"
	"log"

	"github.com/tiansibatuara/golang-url-shortener/models"
	"github.com/tiansibatuara/golang-url-shortener/repository"
	"github.com/tiansibatuara/golang-url-shortener/utils"
	"gorm.io/gorm"
)

type UrlServiceImpl struct {
	UrlRepository repository.UrlRepository
}

func NewUrlServiceImpl(urlRepository repository.UrlRepository) UrlService {
	return &UrlServiceImpl{UrlRepository: urlRepository}
}

// CreateUrl implements UrlService.
func (s *UrlServiceImpl) CreateUrl(originalUrl string) (*models.Url, error) {
	if !utils.IsValidURL(originalUrl) {
		return nil, errors.New("invalid URL")
	}

	shortCode, err := utils.GenerateShortCode()
	if err != nil {
		return nil, errors.New("failed to generate short code")
	}

	existing, err := s.UrlRepository.FindByShortCode(shortCode)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existing != nil {
		return nil, errors.New("short code already exists")
	}

	newUrl := &models.Url{
		Url:       originalUrl,
		ShortCode: shortCode,
	}
	if err := s.UrlRepository.Create(newUrl); err != nil {
		return nil, err
	}
	return newUrl, nil
}

// GetOriginalUrl implements UrlService.
func (s *UrlServiceImpl) GetOriginalUrl(code string) (*models.Url, error) {
	url, err := s.UrlRepository.FindByShortCode(code)
	if err != nil {
		return nil, err
	}
	log.Println(url)
	if err := s.UrlRepository.IncrementAccessCount(code); err != nil {
		return nil, err
	}

	return url, nil
}

// UpdateUrl implements UrlService.
func (s *UrlServiceImpl) UpdateUrl(code string, newUrl string) (*models.Url, error) {
	if !utils.IsValidURL(newUrl) {
		return nil, errors.New("invalid URL")
	}

	url, err := s.UrlRepository.FindByShortCode(code)
	if err != nil {
		return nil, err
	}

	url.Url = newUrl
	if err := s.UrlRepository.Update(url); err != nil {
		return nil, err
	}
	return url, nil
}

// DeleteUrl implements UrlService.
func (s *UrlServiceImpl) DeleteUrl(code string) error {
	return s.UrlRepository.Delete(code)
}

// GetStats implements UrlService.
func (s *UrlServiceImpl) GetStats(code string) (*models.Url, error) {
	return s.UrlRepository.FindByShortCode(code)
}
