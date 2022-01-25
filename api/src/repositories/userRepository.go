package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

//Create a new repository of users
func NewRepositoryOfUsers(db *sql.DB) *users {
	return &users{db}
}

//Create inserts a user into the database
func (u users) Create(user models.User) (uint64, error) {
	statement, erro := u.db.Prepare("INSERT INTO Users (name, nick, email, password) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	ID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ID), nil
}
