package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=tiansibatuara password=dian123 dbname=url_shortener port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}