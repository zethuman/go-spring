package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDB() {
	connStr := "user=postgres dbname=postgres password=admin sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("an error is occured while connecting: %s", err.Error())
	}
	if errors := db.AutoMigrate(&Courses{}).Error; errors != nil {
		log.Fatalf("an error is occured while migrating: %s", err.Error())
	}
	DB = db
}
