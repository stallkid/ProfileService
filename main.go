package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/stallkid/ProfileService/config"
	. "github.com/stallkid/ProfileService/config/dao"
	router "github.com/stallkid/ProfileService/router"

	"github.com/gorilla/mux"
)

var profileDao = ProfileDAO{}
var componentDao = ComponentDAO{}
var config = Config{}

func init() {
	config.Read()
	profileDao.Server = config.Server
	profileDao.Database = config.Database

	componentDao.Server = config.Server
	componentDao.Database = config.Database

	profileDao.Connect()
	componentDao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/profile", router.GetAllComponents).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id}", router.GetComponentByID).Methods("GET")
	r.HandleFunc("/api/v1/profile", router.CreateComponents).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id}", router.UpdateComponents).Methods("PUT")
	r.HandleFunc("/api/v1/profile/{id}", router.DeleteComponents).Methods("DELETE")

	r.HandleFunc("/api/v1/component", router.GetAllComponents).Methods("GET")
	r.HandleFunc("/api/v1/component/{id}", router.GetComponentByID).Methods("GET")
	r.HandleFunc("/api/v1/component", router.CreateComponents).Methods("POST")
	r.HandleFunc("/api/v1/component/{id}", router.UpdateComponents).Methods("PUT")
	r.HandleFunc("/api/v1/component/{id}", router.DeleteComponents).Methods("DELETE")

	var port = ":8091"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
