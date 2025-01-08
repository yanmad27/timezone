package book

import (
	"encore.app/db"
	"github.com/submodule-org/submodule.go/v2"
	"gorm.io/gorm"
)

//encore:service
type Service struct {
	db *gorm.DB
}

func initService() (*Service, error) {
	db, err := db.DBMod.SafeResolve()
	if err != nil {
		return nil, err
	}
	return &Service{db: db}, nil
}

var ServiceMod = submodule.Make[*Service](initService)
