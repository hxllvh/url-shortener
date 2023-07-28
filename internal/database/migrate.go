package database

import (
	"github.com/scetle/urlshortener/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
  err := db.AutoMigrate(&models.URL{})
  if err != nil {
    return err 
  }

  return nil
}
