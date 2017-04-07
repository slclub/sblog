package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Open() (*sql.DB, error) {
	DB, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/sblog?charset=utf8")
	if err != nil {
		panic("Connection mysql error:" + err.Error())
	}

	return DB, nil
}
