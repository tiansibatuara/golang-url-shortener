package repository

import (
	"github.com/tiansibatuara/golang-url-shortener/models"
	"gorm.io/gorm"
)

type UrlRepositoryImpl struct {
	Db *gorm.DB
}

func NewUrlRepository(Db *gorm.DB) UrlRepository {
	return &UrlRepositoryImpl{Db: Db}
}


// Create implements UrlRepository.
func (r *UrlRepositoryImpl) Create(url *models.Url) error {
	res := r.Db.Create(url)
	return res.Error
}

// FindByShortCode implements UrlRepository.
func (r *UrlRepositoryImpl) FindByShortCode(code string) (*models.Url, error) {
	var url models.Url
	res := r.Db.Where("short_code = ?", code).First(&url)
	return &url, res.Error
}

// Update implements UrlRepository.
func (r *UrlRepositoryImpl) Update(url *models.Url) error {
	res := r.Db.Save(url)
	return res.Error
}

// Delete implements UrlRepository.
func (r *UrlRepositoryImpl) Delete(code string) error {
	res := r.Db.Where("short_code = ?", code).Delete(&models.Url{})
	return res.Error
}


// IncrementAccessCount implements UrlRepository.
func (r *UrlRepositoryImpl) IncrementAccessCount(code string) error {
	res := r.Db.Where("short_code = ?", code).
		Update("access_count", gorm.Expr("access_count + 1"))
	return res.Error
}
