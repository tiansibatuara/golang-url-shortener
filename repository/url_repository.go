package repository

import (
	"errors"
	"url_shortener/models"

	"gorm.io/gorm"
)

type URLRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(url *models.Url) error {
	return r.db.Create(url).Error
}

func (r *URLRepository) FindByShortCode(code string) (*models.Url, error) {
	var url models.Url
	res := r.db.Where("short_code = ?", code).First(&url)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &url, res.Error
}

func (r *URLRepository) Update(url *models.Url) error {
	return r.db.Save(url).Error
}

func (r *URLRepository) Delete(url *models.Url) error {
	return r.db.Delete(url).Error
}

func (r *URLRepository) IncrementAccessCount(code string) error {
	return r.db.Model(&models.Url{}).
		Where("short_code = ?", code).
		Update("access_count", gorm.Expr("access_count + 1")).
		Error
}
