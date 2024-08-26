// driver/db.go
package driver

import (
	"github.com/devopscorner/golang-bedrock/src/config"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDatabase() {
	if config.DbConnection() == "mysql" {
		ConnectMySQL()
		DB = DB_MySQL
	} else if config.DbConnection() == "postgres" {
		ConnectPSQL()
		DB = DB_PSQL
	} else {
		ConnectSQLite()
		DB = DB_SQLite
	}
}
