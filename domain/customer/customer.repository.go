package domain

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
