package main

func (a *App) handleRequests() {
	a.mux.HandleFunc("/", homePage).Methods("GET")
	a.mux.HandleFunc("/address/{postal_code}", a.singleAddress).Methods("GET")
}
