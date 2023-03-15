package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	domain "go-crud/domain/customer"
	service "go-crud/service/customer"
	"net/http"
)

type CustomerHandlers struct {
	Service service.CustomerService
}

func responseWrapper(w http.ResponseWriter, r *http.Request, payload *[]domain.Customer) {
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(payload)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	}
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.Service.GetAll()
	if err != nil {
		fmt.Printf("Error during get all customers: %s", err.Error())
	}
	responseWrapper(w, r, &customers)
}

//func GetCustomer(w http.ResponseWriter, r *http.Request) {
//	customers := getAllCustomers()
//	reqVars := mux.Vars(r)
//	id := reqVars["id"]
//	var customer []domain.Customer
//
//	for _, v := range customers {
//		if v.Id == id {
//			customer = append(customer, v)
//			break
//		}
//	}
//
//	responseWrapper(w, r, &customer)
//}
//
//func CreateCustomer(w http.ResponseWriter, r *http.Request) {
//	result := make(map[string]interface{})
//	result["success"] = true
//	w.Header().Add("Content-Type", "applications/json")
//	json.NewEncoder(w).Encode(result)
//}
