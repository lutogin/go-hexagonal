package customers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
)

func getAllCustomers() []Customer {
	return []Customer{
		{"1", "Tom", "Brinska 22", "1123"},
		{"2", "Derek", "NY 15 ave", "342"},
		{"3", "Dan", "LA main str 18", "24563"},
	}
}

func responseWrapper(w http.ResponseWriter, r *http.Request, payload *[]Customer) {
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(payload)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	}
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := getAllCustomers()
	responseWrapper(w, r, &customers)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	customers := getAllCustomers()
	reqVars := mux.Vars(r)
	id := reqVars["id"]
	var customer []Customer

	for _, v := range customers {
		if v.id == id {
			customer = append(customer, v)
			break
		}
	}

	responseWrapper(w, r, &customer)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]interface{})
	result["success"] = true
	w.Header().Add("Content-Type", "applications/json")
	json.NewEncoder(w).Encode(result)
}
