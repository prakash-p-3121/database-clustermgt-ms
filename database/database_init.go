package database

import "database/sql"

var databaseConnection *sql.DB

func SetDatabaseInstance(databaseInst *sql.DB) {
	databaseConnection = databaseInst
}

func GetDatabaseInstance() *sql.DB {
	return databaseConnection
}
