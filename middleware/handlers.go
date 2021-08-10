package middleware

import (
	"encoding/json"
	"fmt"
	"gobackend/controllers"
	"gobackend/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type conn interface {
}

func gormConnection() *gorm.DB {
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

func CreateBeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var beer models.Beer

	err := json.NewDecoder(r.Body).Decode(&beer)

	if err != nil {
		log.Fatalf("parsing problems, %v", err)
	}
	created := controllers.InsertBeer(beer)
	res, _ := json.Marshal(created)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "pkglication/json")
	params := mux.Vars(r)
	var beers []models.Beer
	if len(params) == 0 {
		beers = controllers.FindBeer()
	} else {
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			log.Fatalf("Unable to convert the string into int.  %v", err)
		}

		beers = controllers.FindBeer(int64(id))

		if err != nil {
			log.Fatalf("Unable to get beer. %v", err)
		}
	}
	res, _ := json.Marshal(beers)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "pkglication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var beer models.Beer

	err = json.NewDecoder(r.Body).Decode(&beer)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := controllers.UpdateBeer(int64(id), beer)
	msg := fmt.Sprintf("Beer updated successfully. Total rows/record affected %v", updatedRows)

	res, _ := json.Marshal(response{
		ID:      int64(id),
		Message: msg,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBeer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := controllers.DeleteBeer(int64(id))

	msg := fmt.Sprintf("Beer removed successfully. Total rows/record affected %v", deletedRows)

	res, _ := json.Marshal(response{
		ID:      int64(id),
		Message: msg,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
