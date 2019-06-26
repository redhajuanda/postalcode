package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	db  *sql.DB
	mux *mux.Router
}

func newApp(db *sql.DB, mux *mux.Router) *App {
	a := App{db, mux}
	a.handleRequests()
	return &a
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
