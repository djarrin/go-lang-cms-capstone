package routeHandlers

import (
	d "go-lang-cms-capstone/dataHandler"
	"net/http"
)

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
	if newPhone := newCustomerData.Phone; newPhone != 0 {
		existingCustomer.Phone = newPhone
	}
	//Contacted if omitted will always be set to false or if the field is actually set to false, it will be set to true if field is set to true
	existingCustomer.Contacted = newCustomerData.Contacted

	return existingCustomer
}
