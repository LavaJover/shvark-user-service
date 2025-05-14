package postgres

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustInitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatalf("failed to connect to db: %v", err)
	}

	db.AutoMigrate(&UserModel{})

	return db
}