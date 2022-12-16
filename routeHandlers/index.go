package routeHandlers

import (
	"encoding/json"
	d "go-lang-cms-capstone/dataHandler"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/docs.html")
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var returnSlice = []d.CustomerReturn{}

	for id, customer := range d.Data {
		returnSlice = append(returnSlice, customerReturnObject(id, customer))
	}

	json.NewEncoder(w).Encode(returnSlice)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		defer handleExit(w)
	}

	_, isPresent := d.Data[intId]

	if isPresent {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customerReturnObject(intId, d.Data[intId]))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func customerReturnObject(id int, c d.Customer) d.CustomerReturn {
	return d.CustomerReturn{
		Id:        id,
		Name:      c.Name,
		Role:      c.Role,
		Email:     c.Email,
		Phone:     c.Phone,
		Contacted: c.Contacted,
	}
}

func handleExit(w http.ResponseWriter) {
	if r := recover(); r != nil {
		if he, ok := r.(d.HttpError); ok {
			http.Error(w, he.Message, he.Status)
		} else {
			panic(r)
		}
	}
}

func exit(status int, message string) {
	panic(d.HttpError{Status: status, Message: message})
}
