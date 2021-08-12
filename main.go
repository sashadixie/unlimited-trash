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
	fmt.Println("Starting server on the port...")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
