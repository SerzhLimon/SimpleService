package main

import (
	database "SimpleService/pkg/db"
	"SimpleService/pkg/transport"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const limit = 3
const db_options = "host=localhost port=9300 user=illustrv dbname=postgres sslmode=disable"

func main() {

	sql, err := database.New(db_options)
	if err != nil {
		log.Fatalln(err)
	}

	server := transport.New(sql, limit)
	fmt.Println("Server started ... ")

	mux := http.NewServeMux()
	mux.HandleFunc("/", server.SearchHandler)
	mux.HandleFunc("/data", server.SearchHandlerById)
	mux.HandleFunc("/create", server.PublishHandler)
	mux.HandleFunc("/admin", server.AuthorizationHandler)
	mux.HandleFunc("/connect", server.CheckDataHandler)
	err = http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatal(err)
	}
}


//http://localhost:8888/?page=1 - home page