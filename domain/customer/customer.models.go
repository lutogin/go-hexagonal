package domain

import "time"

type Customer struct {
	Id          string    `json:"id" xml:"id"`
	Name        string    `json:"name" xml:"name"`
	Address     string    `json:"address" xml:"address"`
	Zip         string    `json:"zip" xml:"zip"`
	DateOfBirth time.Time `json:"dateOfBirth" xml:"dateOfBirth"`
}
