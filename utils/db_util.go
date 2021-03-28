package util

import (
	"database/sql"
	"os"
)

func PanicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
func DBConn() *sql.DB {
	connection_string := os.Getenv("MYSQL_CONNECTION_STRING")
	db, err := sql.Open("mysql", connection_string)
	PanicError(err)

	return db
}
