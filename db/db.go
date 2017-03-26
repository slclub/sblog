package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	DB, err := sql.Open("mysql", "root:@127.0.0.1/sblog?charset=utf8")
	if err != nil {
		panic("Connection mysql error:" + err.Error())
	}

	return DB
}
