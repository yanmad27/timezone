package book

import (
	"encore.app/db"
	"gorm.io/gorm"
)

//encore:service
type Service struct {
	db *gorm.DB
}

func initService() (*Service, error) {
	db := db.Init()
	return &Service{db: db}, nil
}
