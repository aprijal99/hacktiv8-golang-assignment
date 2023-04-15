package db

import (
	"fmt"
	"mygram/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "subang12345"
	port     = 5432
	dbname   = "mygram"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{}, &gorm.Config{TranslateError: true})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(entity.User{}, entity.Photo{}, entity.SocialMedia{}, entity.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
