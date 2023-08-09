package models

import (
  "gorm.io/gorm"
)

// for DB migrations
type URL struct {
  gorm.Model
  OriginalURL string
  ShortURL string
}

// for templates data transfer
type DataURL struct {
  gorm.Model
  OriginalURL string
  ShortURL    string
  ExistingURL string
}
