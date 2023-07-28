package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MyDB struct {
  db *gorm.DB
}

func NewDB() (*MyDB, error) {
  dsn := "your dsn here"
  dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }
  return &MyDB{db: dbConnection}, nil
}

func (d *MyDB) GetDB() *gorm.DB {
  return d.db
}
