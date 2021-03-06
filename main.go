package main

import (
	"fmt"
	"log"
	"net/http"

	"example/test/handler"

	"github.com/gorilla/mux"
)

func main() {
	// db.Init()
	fmt.Printf("start")
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/companies").HandlerFunc(handler.GetCompanies)
	sub.Methods("POST").Path("/companies").HandlerFunc(handler.SaveCompany)
	sub.Methods("GET").Path("/companies/{name}").HandlerFunc(handler.GetCompany)
	sub.Methods("PUT").Path("/companies/{name}").HandlerFunc(handler.UpdateCompany)
	sub.Methods("DELETE").Path("/companies/{name}").HandlerFunc(handler.DeleteCompany)

	log.Fatal(http.ListenAndServe(":3000", router))

}
