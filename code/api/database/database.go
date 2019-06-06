package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Connect(dialect, uri string) error {
	var err error
	db, err = gorm.Open(dialect, uri)
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	db.Close()
}

func Migrate() {
	db.AutoMigrate(&User{})
}
