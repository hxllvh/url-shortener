package models

import (
  "gorm.io/gorm"
)

type URL struct {
  gorm.Model
  OriginalURL string
  ShortURL    string
}
