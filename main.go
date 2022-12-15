package main

import (
	"fmt"
	"go-lang-cms-capstone/routeHandlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", routeHandlers.ServeIndex).Methods("GET")

	fmt.Println("Server is starting on port 3000...")
	// Pass the customer router into ListenAndServe
	http.ListenAndServe(":3000", router)
}
