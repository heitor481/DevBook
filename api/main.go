package main

import (
	"api/src/router"
	"log"
	"net/http"
)

func main() {
	println("Enter point")

	router := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(":5000", router))
}
