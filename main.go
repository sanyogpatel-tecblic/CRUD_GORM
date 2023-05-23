package main

import (
	"log"

	"github.com/sanyogpatel-tecblic/CRUD_GORM/controller/models"
	"github.com/sanyogpatel-tecblic/CRUD_GORM/controller/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres dbname=student password=root port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migration completed successfully")
	routes.SetupRouter(db)

}
