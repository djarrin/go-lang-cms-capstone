package main

import (
	"fmt"
	rh "go-lang-cms-capstone/routeHandlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", rh.ServeIndex).Methods("GET")
	router.HandleFunc("/customers", rh.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", rh.GetCustomer).Methods("GET")
	router.HandleFunc("/customers", rh.AddCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", rh.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", rh.DeleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
	return
}
