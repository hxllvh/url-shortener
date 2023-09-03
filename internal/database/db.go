package database

import (
	"github.com/scetle/url-shortener/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type myDB struct {
  *gorm.DB
}

func NewDB() (*myDB, error) {
  dsn := "host=localhost user=postgres dbname=urlshortener port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }
  return &myDB{db}, nil
}

func (db *myDB) AddURL(newURL models.URL) (string, error) {
  var url models.URL
  db.DB.Where("original_url = ?", newURL.OriginalURL).First(&url)
  if url.OriginalURL != "" {
    return url.OriginalURL, nil
  }

  result := db.Create(&newURL)
  return url.OriginalURL, result.Error
}
