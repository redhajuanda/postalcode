package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) singleAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["postal_code"]
	db := a.db

	var postal_code PostalCode
	var arr_postal_code []PostalCode

	query := fmt.Sprintf("SELECT urban, sub_district, city, province_code, postal_code FROM db_postal_code_data WHERE postal_code=%v LIMIT 1", id)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
		response := Response{
			Status:  400,
			Message: "Bad Request",
			Data:    nil,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	if rows.Next() {
		if err := rows.Scan(&postal_code.Urban, &postal_code.SubDistrict, &postal_code.City, &postal_code.Province.ProvinceCode, &postal_code.PostalCode); err != nil {
			log.Fatal(err.Error())
		}
	} else {
		response := Response{
			Status:  400,
			Message: "No Data Found",
			Data:    nil,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	query = fmt.Sprintf("SELECT province_code, province_name FROM db_province_data WHERE province_code=%v", postal_code.Province.ProvinceCode)

	rows, err = db.Query(query)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(&postal_code.Province.ProvinceCode, &postal_code.Province.ProvinceName); err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Print(postal_code)
			arr_postal_code = append(arr_postal_code, postal_code)
		}
	}
	response := Response{
		Status:  200,
		Message: "Success",
		Data:    arr_postal_code,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func allAddress(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var postal_code PostalCode
	var arr_postal_code []PostalCode

	query := fmt.Sprintf("SELECT urban, sub_district, city, province_code, postal_code FROM db_postal_code_data")

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(&postal_code.Urban, &postal_code.SubDistrict, &postal_code.City, &postal_code.Province.ProvinceCode, &postal_code.PostalCode); err != nil {
			log.Fatal(err.Error())
		} else {
			// fmt.Println(postal_code)
			query = fmt.Sprintf("SELECT province_code, province_name FROM db_province_data WHERE province_code=%v", postal_code.Province.ProvinceCode)

			rows2, err := db.Query(query)
			if err != nil {
				log.Print(err)
			}
			for rows2.Next() {
				if err := rows2.Scan(&postal_code.Province.ProvinceCode, &postal_code.Province.ProvinceName); err != nil {
					log.Fatal(err.Error())
				} else {
					// fmt.Print(postal_code)
					arr_postal_code = append(arr_postal_code, postal_code)
				}
			}
		}
	}

	response := Response{
		Status:  200,
		Message: "Success",
		Data:    arr_postal_code,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
	fmt.Println("GET '/' Request")
}
