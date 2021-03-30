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
	connection_string := "doadmin:clhx24ycdanxqvr3@tcp(db-mysql-nyc3-67043-do-user-8989383-0.b.db.ondigitalocean.com:25060)/defaultdb"
	db, err := sql.Open("mysql", connection_string)
	PanicError(err)

	return db
}
