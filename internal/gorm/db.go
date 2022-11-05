package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(conn string) error {
	database, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = database
	err = Migrate()
	return err
}

func Migrate() error {
	// todo add migration structures
	err := GetDB().AutoMigrate()
	return err
}

func GetDB() *gorm.DB {
	return db
}
