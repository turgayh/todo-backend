package util

import (
	"database/sql"
)

func PanicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
func DBConn() *sql.DB {
	db, err := sql.Open("mysql", "root:helloworld@tcp(127.0.0.1:6603)/todo")
	PanicError(err)

	return db
}
