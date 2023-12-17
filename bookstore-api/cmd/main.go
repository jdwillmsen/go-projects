package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jdwillmsen/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	port := "9404"
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
