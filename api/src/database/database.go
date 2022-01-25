package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb" //DRiVER
)

//Connect connecto to a database and return
func Connect() (*sql.DB, error) {
	db, erro := sql.Open("sqlserver", config.ConnectionString)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
