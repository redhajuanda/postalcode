package main

import (
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

func main() {

	db := connect()
	mux := mux.NewRouter()
	app := newApp(db, mux)
	http.ListenAndServe(":3000", app.mux)
}
