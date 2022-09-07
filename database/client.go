package database

import (
	"go-restful-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var dbError error

func Connection(connectingString string) {
	Instance, dbError = gorm.Open(postgres.Open(connectingString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("can not connect to Databases")
	}
	log.Println("Connected to Database")
}
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed")
}
