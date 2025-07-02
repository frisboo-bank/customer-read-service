package models

import (
	"time"

	cid "frisboo-bank/pkg/core/cid"
)

type CustomerContactInformation struct {
	Emails             []Email
	HomePhoneNumbers   []PhoneNumber
	MobilePhoneNumbers []PhoneNumber
	WorkPhoneNumbers   []PhoneNumber
}

type CustomerIdentificationDocuments struct{}

type CustomerPersonalInformations struct {
	Title                string
	FirstName            string
	MiddleName           string
	LastName             string
	FirstNameInEnglish   string
	MiddleNameInEnglish  string
	LastNameInEnglish    string
	DateOfBirth          time.Time
	CountryOfBirthId     string
	NationalitiesId      []string
	GenderId             string
	IsPoliticallyExposed bool
	IsUSPerson           bool
	MaritalStatusId      string
	NumberOfDependents   uint8
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type CustomerAddresses struct {
	CorrespondanceAddress []Address
	ResidentialAddress    []Address
	WorkAddress           []Address
}

type Customer struct {
	Cid                     cid.Cid
	ContactInformation      CustomerContactInformation
	IdentificationDocuments CustomerIdentificationDocuments
	PersonalInformation     CustomerPersonalInformations
	Addresses               CustomerAddresses
	CreatedAt               time.Time
}
