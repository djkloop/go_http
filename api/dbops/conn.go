package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/GoHttp?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
