package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//The connection string
	ConnectionString = ""
	//The port to connect
	Port = 0
)

//Charge will initialize the env variables
func Charge() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}

	ConnectionString = fmt.Sprintf("server=%s;user id=AVELL;database=%s",
		os.Getenv("DB_SERVER"),
		os.Getenv("DB_NAME"))
}
