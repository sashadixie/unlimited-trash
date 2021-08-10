package router

import (
	"encoding/json"
	"gobackend/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type AppResponse struct {
	Message string `json:"message,omitempty"`
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	res := AppResponse{
		Message: "Looks like we are lost",
	}
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(res)
}

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/beer/{id}", middleware.GetBeer).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/beer", middleware.GetBeer).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/beer/create", middleware.CreateBeer).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/beer/{id}", middleware.UpdateBeer).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete/{id}", middleware.DeleteBeer).Methods("DELETE", "OPTIONS")
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	return router
}
