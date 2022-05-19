package mysql

import (
	"database/sql"

	"github.com/Dramane-dev/todolist-api/api/service"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDatabase struct {
	connection *gorm.DB
}

var _ service.UserService = &MySQLDatabase{}
var _ service.ProjectService = &MySQLDatabase{}

func New(DatabaseDriver, DatabaseUser, DatabasePassword, DatabasePort, DatabaseHost, DatabaseName string) *MySQLDatabase {
	// databaseUrl := DatabaseUser + ":" + DatabasePassword + "@/" + DatabaseName
	databaseUrl := DatabaseUser + ":" + DatabasePassword + "@tcp(" + DatabaseHost + ":+" + DatabasePort + ")/" + DatabaseName + "?charset=utf8&parseTime=True&loc=Local"
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

// func MigrateInterfaces(sqlGormDatabase *gorm.DB, user *models.User, project *models.Project) {
// 	errWhenAutoMigrateType := sqlGormDatabase.AutoMigrate(&models.User{}, &models.Project{}).Error()

// 	if len(errWhenAutoMigrateType) > 0 {
// 		panic(errWhenAutoMigrateType)
// 	}
// }
