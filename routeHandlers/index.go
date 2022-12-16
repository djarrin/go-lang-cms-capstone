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
		return
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

	newID := d.GetNextAvailaleID()
	customer := &d.Customer{}
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(customer); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		defer handleExit(w)
		return
	}

	if customer.Email == "" || customer.Name == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		resp := make(map[string]string)
		resp["message"] = "Neither Email or Name fields may be empty"
		jsonResp, _ := json.Marshal(resp)

		w.Write(jsonResp)
		defer handleExit(w)
		return
	}

	d.Data[newID] = *customer

	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(customerReturnObject(newID, d.Data[newID]))
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		defer handleExit(w)
		return
	}

	customerToUpdate, isPresent := d.Data[intId]

	if !isPresent {
		w.WriteHeader(http.StatusNotFound)
		defer handleExit(w)
		return
	}

	customer := &d.Customer{}
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(customer); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		defer handleExit(w)
		return
	}

	d.Data[intId] = updateCustomerStruct(customerToUpdate, *customer)

	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(customerReturnObject(intId, d.Data[intId]))
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		defer handleExit(w)
		return
	}

	_, isPresent := d.Data[intId]

	if isPresent {
		delete(d.Data, intId)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

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

func updateCustomerStruct(existingCustomer d.Customer, newCustomerData d.Customer) d.Customer {
	if newName := newCustomerData.Name; newName != "" {
		existingCustomer.Name = newName
	}
	if newRole := newCustomerData.Role; newRole != "" {
		existingCustomer.Role = newRole
	}
	if newEmail := newCustomerData.Email; newEmail != "" {
		existingCustomer.Email = newEmail
	}
	if newPhone := newCustomerData.Phone; newPhone != "" {
		existingCustomer.Phone = newPhone
	}
	//Contacted if omitted will always be set to false or if the field is actually set to false, it will be set to true if field is set to true
	existingCustomer.Contacted = newCustomerData.Contacted

	return existingCustomer
}
