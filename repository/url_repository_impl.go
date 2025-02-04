package repository

import (
	"errors"
	"log"

	"github.com/tiansibatuara/golang-url-shortener/models"
	"gorm.io/gorm"
)

type UrlRepositoryImpl struct {
	db *gorm.DB
}

func NewUrlRepository(Db *gorm.DB) UrlRepository {
	return &UrlRepositoryImpl{db: Db}
}

// Create implements UrlRepository.
func (r *UrlRepositoryImpl) Create(url *models.Url) error {
	res := r.db.Create(url)
	return res.Error
}

// FindByShortCode implements UrlRepository.
func (r *UrlRepositoryImpl) FindByShortCode(code string) (*models.Url, error) {
	var url models.Url
	res := r.db.Where("short_code = ?", code).First(&url)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil // No record found
		}
		return nil, res.Error // Database error
	}
	return &url, nil
}

// Update implements UrlRepository.
func (r *UrlRepositoryImpl) Update(url *models.Url) error {
	res := r.db.Save(url)
	return res.Error
}

// Delete implements UrlRepository.
func (r *UrlRepositoryImpl) Delete(code string) error {
	res := r.db.Where("short_code = ?", code).Delete(&models.Url{})
	return res.Error
}

// IncrementAccessCount implements UrlRepository.
func (r *UrlRepositoryImpl) IncrementAccessCount(code string) error {
	log.Println("Incrementing access count for short code:", code)
	res := r.db.Model(&models.Url{}).Where("short_code = ?", code).Update("access_count", gorm.Expr("access_count + 1"))
	if res.Error != nil {
		log.Println("Error incrementing access count:", res.Error)
	}
	return res.Error
}
