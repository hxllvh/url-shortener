package database

import (
	"github.com/scetle/url-shortener/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
  err := db.AutoMigrate(&models.URL{})
  if err != nil {
    return err 
  }

  return nil
}
