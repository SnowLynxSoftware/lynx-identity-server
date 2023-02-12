package database

import (
	"SnowLynxSoftware/lynx-identity-server/pkg/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DBConnection *sql.DB
)

func GetDBConnection() *sql.DB {

	db, err := sql.Open("mysql", config.AppConfig.DBConnectionString)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
}
