package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/huckflinn/go-bookstore/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
