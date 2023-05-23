package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres dbname=test password=secret port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	models.User

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migration completed successfully")

}
