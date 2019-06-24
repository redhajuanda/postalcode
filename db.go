package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "redhajuanda:4m4l4n@tcp(app.amalan.com)/kodepos")

	if err != nil {
		log.Panic(err.Error())
	}
	return db
}
