package controllers

import "net/http"

//Create user
func CreateUser(rw http.ResponseWriter, response *http.Request) {
	rw.Write([]byte("Creating user"))
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
