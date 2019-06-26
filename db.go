package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redhajuanda/postalcode/config"
)

func connect() *sql.DB {
	conf := config.New()

	db_string := fmt.Sprintf("%v:%v@tcp(%v)/%v", conf.DB.Username, conf.DB.Password, conf.DB.Host, conf.DB.Database)
	db, err := sql.Open("mysql", db_string)

	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println(db_string)
	// fmt.Println("CONNNNNNN")
	return db
}
