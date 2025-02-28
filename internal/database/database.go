package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB(connStr string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	return db
}
