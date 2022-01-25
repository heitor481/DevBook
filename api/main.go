package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Charge()

	fmt.Print(config.ConnectionString)
	router := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
