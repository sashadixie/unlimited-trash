package controllers

import (
	"fmt"
	"gobackend/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func createConnection() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")

	return db
}

func InsertBeer(beer models.Beer) models.Beer {
	db := createConnection()
	db.Create(&beer)
	return beer
}

func UpdateBeer(id int64, beer models.Beer) int64 {
	db := createConnection()
	return db.Model(&beer).Where("id = ?", id).Updates(beer).RowsAffected
}

func DeleteBeer(id int64) int64 {
	var beer models.Beer
	db := createConnection()
	return db.Where("id = ?", id).Delete(beer).RowsAffected
}

func FindBeer(id ...int64) []models.Beer {
	db := createConnection()
	var Beers []models.Beer
	if len(id) > 0 {
		db.Where("id = ?", id).Find(&Beers)
	} else {
		db.Find(&Beers)
	}
	return Beers
}
