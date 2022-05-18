package mysql

import (
	"database/sql"

	"github.com/Dramane-dev/todolist-api/api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDatabase struct {
	connection *gorm.DB
}

var _ config.UserService = &MySQLDatabase{}
var _ config.ProjectService = &MySQLDatabase{}

func New(DatabaseDriver, DatabaseUser, DatabasePassword, DatabasePort, DatabaseHost, DatabaseName string) *MySQLDatabase {
	databaseUrl := DatabaseUser + ":" + DatabasePassword + "@/" + DatabaseName
	sqlDatabase, sqlConnectionError := sql.Open(DatabaseDriver, databaseUrl)

	if sqlConnectionError != nil {
		panic(sqlConnectionError)
	}

	sqlGormDatabase, sqlGormConnectionError := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDatabase,
	}), &gorm.Config{})

	if sqlGormConnectionError != nil {
		panic(sqlGormConnectionError)
	}

	return &MySQLDatabase{
		connection: sqlGormDatabase,
	}
}
