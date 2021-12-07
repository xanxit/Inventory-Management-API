package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"xanxit.com/helper"
	"xanxit.com/controllers"
)

//Connection mongoDB with helper class
var collection = helper.ConnectDB()

func main() {
	//Init Router
	r := mux.NewRouter()

	// arrange our route
	r.HandleFunc("/api/books", controllers.getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.getBook).Methods("GET")
	r.HandleFunc("/api/books", controllers.createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", controllers.updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controllers.deleteBook).Methods("DELETE")

	// set our port address
	log.Fatal(http.ListenAndServe(":1027", r))

}
