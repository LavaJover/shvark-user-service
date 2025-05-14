package postgres

import (
	"log"

	"github.com/LavaJover/shvark-user-service/internal/infrastructure/postgres/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustInitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatalf("failed to connect to db: %v", err)
	}

	db.AutoMigrate(&model.UserModel{})

	return db
}