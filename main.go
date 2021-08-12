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
	fmt.Println("Port from os..." + port)
	if port == "" {
		port = "9000"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}
