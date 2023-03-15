package domain

import "time"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Tom", "Brinska 22", "1123", time.Date(1991, 02, 12, 0, 0, 0, 0, time.UTC)},
		{"2", "Derek", "NY 15 ave", "342", time.Date(1951, 06, 16, 0, 0, 0, 0, time.UTC)},
		{"3", "Dan", "LA main str 18", "24563", time.Date(1998, 12, 12, 0, 0, 0, 0, time.UTC)},
	}

	return CustomerRepositoryStub{customers}
}
