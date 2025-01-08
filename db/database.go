package db

import (
	"log"
	"time"

	"github.com/submodule-org/submodule.go/v2"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	time.Local = loc

	dbURL := "postgres://postgres@localhost:5432"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Book{})
	return db
}

var DBMod = submodule.Make[*gorm.DB](Init)
