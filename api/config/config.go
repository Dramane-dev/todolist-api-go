package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB(DatabaseDriver, DatabaseUser, DatabasePassword, DatabasePort, DatabaseHost, DatabaseName string) (db *sql.DB, err error) {
	databaseUrl := DatabaseUser + ":" + DatabasePassword + "@/" + DatabaseName
	db, err = sql.Open(DatabaseDriver, databaseUrl)
	return
}
