package main

import (
	"log"
	"net/http"

	"github.com/bijan/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe("localhost:9010",router))
}
