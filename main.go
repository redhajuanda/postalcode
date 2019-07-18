package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

type App struct {
	db  *sql.DB
	mux *mux.Router
}

func main() {
	db := connect()
	mux := mux.NewRouter()
	app := newApp(db, mux)
	err := http.ListenAndServe(":5000", app.mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func newApp(db *sql.DB, mux *mux.Router) *App {
	a := App{db, mux}
	a.handleRequests()
	return &a
}

// func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	a.mux.ServeHTTP(w, r)
// }
