package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"app/helper"
	"log"
)

func GetWeapons(w http.ResponseWriter, r *http.Request) {
	weapons := helper.LoadWeaponsCsv()
	json.NewEncoder(w).Encode(weapons)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/weapons", GetWeapons).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}