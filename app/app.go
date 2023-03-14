package app

import (
	"fmt"
	"go-crud/src/api"
	"go-crud/src/customers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is running...")
}

func Start() {
	rounter := mux.NewRouter()

	rounter.HandleFunc("/customers", customers.GetAllCustomers).Methods(http.MethodGet)
	rounter.HandleFunc("/customers/{id:[0-9]+}", customers.GetCustomer).Methods(http.MethodGet)
	rounter.HandleFunc("/customers", customers.CreateCustomer).Methods(http.MethodPost)
	rounter.HandleFunc("/api/time", api.ApiTime).Methods(http.MethodGet)
	rounter.HandleFunc("/", mainPage).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8181", rounter))
}
