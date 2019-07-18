package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redhajuanda/postalcode/config"
)

// var (
// 	DB *sql.DB
// )

func connect() *sql.DB {
	conf := config.New()

	db_string := fmt.Sprintf("%v:%v@tcp(%v)/%v", conf.DB.Username, conf.DB.Password, conf.DB.Host, conf.DB.Database)
	db, err := sql.Open("mysql", db_string)

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(300 * time.Second)

	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println(db_string)
	return db
}

// // OpenDbPool for open DB Pool
// func OpenDbPool() {
// 	DB = connect()
// 	DB.SetMaxOpenConns(100)
// 	DB.SetMaxIdleConns(10)
// 	DB.SetConnMaxLifetime(300 * time.Second)
// }
