package database

import (
	"fmt"
	"log"
	"task/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
var DNS = "host=localhost user=postgres password=admin dbname=Users port=5432 sslmode=disable"

func Migration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal("not connected to the database")
	}
	fmt.Print("connecteed to the database")
	DB.AutoMigrate(&model.User{})
}
