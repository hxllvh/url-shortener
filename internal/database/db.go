package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MyDB struct {
  *gorm.DB
}

func NewDB() (*MyDB, error) {
  dsn := "host=localhost user=postgres dbname=urlshortener port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }
  return &MyDB{db}, nil
}
