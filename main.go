package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/stallkid/ProfileService/config"
	. "github.com/stallkid/ProfileService/config/dao"
	profilerouter "github.com/stallkid/ProfileService/router"

	"github.com/gorilla/mux"
)

var dao = ProfileDAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/profile", profilerouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id}", profilerouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/profile", profilerouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id}", profilerouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/profile/{id}", profilerouter.Delete).Methods("DELETE")

	var port = ":8091"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
