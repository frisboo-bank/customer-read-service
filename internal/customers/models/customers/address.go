package models

import "time"

type Address struct {
	Id           string
	AddressLine1 string
	AddressLine2 string
	AddressLine3 string
	AddressLine4 string
	AddressLine5 string
	State        string
	District     string
	City         string
	Region       string
	PostalCode   string
	CountryID    string
	IsPrimary    bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
