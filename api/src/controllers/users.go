package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Create user
func CreateUser(rw http.ResponseWriter, response *http.Request) {
	requestBody, erro := ioutil.ReadAll(response.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User

	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.NewRepositoryOfUsers(db)
	userID, erro := repository.Create(user)
	if erro != nil {
		log.Fatal(erro)
	}

	rw.Write([]byte(fmt.Sprintf("User inserted: %d", userID)))
}

//Get all Users
func GetAllUsers(rw http.ResponseWriter, response *http.Request) {
	rw.Write([]byte("Get all users"))
}

//Get user
func GetUser(rw http.ResponseWriter, response *http.Request) {
	rw.Write([]byte("Get user"))
}

//Update user
func UpdateUser(rw http.ResponseWriter, response *http.Request) {
	rw.Write([]byte("Update user"))
}

//Delete user
func DeleteUser(rw http.ResponseWriter, response *http.Request) {
	rw.Write([]byte("Delete user"))
}
