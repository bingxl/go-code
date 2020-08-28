package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbConn *sql.DB
var err error

func init() {
	dbConn, err = sql.Open("mysql", "root:lxb106170$$@tcp(localhost:3306)/video?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

}
