package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go-restapi/config"
	"go-restapi/config/dao"
	router "go-restapi/router"
)

const (
	sensoresGetAll = `/v1/inter/sensores` // GET
	sensoresInsert = `/v1/inter/sensores` //POST
)

var dataacess = dao.SensorDao{}
var configuration = config.Config{}

func init() {
	configuration.Read()

	dataacess.Server = configuration.Server
	dataacess.Database = configuration.Database
	dataacess.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc(sensoresGetAll, router.GetAll).Methods("GET")
	r.HandleFunc(sensoresInsert, router.Create).Methods("POST")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}