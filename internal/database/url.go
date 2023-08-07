package database

import (
	"github.com/scetle/urlshortener/internal/models"
	"gorm.io/gorm"
)

func AddURL(db *gorm.DB, newUrl models.URL) error {
  result := db.Create(&newUrl)
  return result.Error
}

func CheckIfExists(db *gorm.DB, original string) (string, error) {
  var url models.URL
  result := db.Where("original_url = ?", original).First(&url)
  if result.Error == nil {
    return url.OriginalURL, nil
  }

  return url.OriginalURL, result.Error
}
