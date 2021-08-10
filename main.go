package main

import (
	"fmt"
	"gobackend/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8082...")

	log.Fatal(http.ListenAndServe(":8082", r))
}
