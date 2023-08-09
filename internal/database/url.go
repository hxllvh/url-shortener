package database

import (
	"github.com/scetle/urlshortener/internal/models"
	"gorm.io/gorm"
)

func AddURL(db *gorm.DB, newURL models.URL) error {
  result := db.Create(&newURL)
  return result.Error
}

func CheckIfExists(db *gorm.DB, original string) string {
  var url models.URL
  db.Where("original_url = ?", original).First(&url)
  return url.OriginalURL
}
