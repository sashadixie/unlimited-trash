package main

import (
	"fmt"
	"gobackend/router"
	"log"
	"net/http"
	"os"
)

func main() {
	r := router.Router()
	port := os.Getenv("PORT")
	fmt.Println("Starting server on the port...")

	log.Fatal(http.ListenAndServe(port, r))
}
