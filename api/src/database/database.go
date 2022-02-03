package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

//Connect connecto to a database and return
func Connect() (*sql.DB, error) {
	db, erro := sql.Open("pgx", config.ConnectionString)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
