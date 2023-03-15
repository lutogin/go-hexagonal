package app

import (
	"go-crud/domain/api"
	domain "go-crud/domain/customer"
	handler "go-crud/handler/customer"
	service "go-crud/service/customer"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	ch := handler.CustomerHandlers{Service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/api/time", api.ApiTime).Methods(http.MethodGet)
	//router.HandleFunc("/customers/{id:[0-9]+}", customer.GetCustomer).Methods(http.MethodGet)
	//router.HandleFunc("/customers", customer.CreateCustomer).Methods(http.MethodPost)
	//router.HandleFunc("/", mainPage).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8181", router))
}
