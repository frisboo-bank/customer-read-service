package customers

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

type CustomerPersonalInformation struct {
	Title                string
	FirstName            string
	MiddleName           string
	LastName             string
	FirstNameInEnglish   string
	MiddleNameInEnglish  string
	LastNameInEnglish    string
	DateOfBirth          time.Time
	CountryOfBirthID     string
	NationalitiesID      []string
	GenderID             string
	IsPoliticallyExposed bool
	IsUSPerson           bool
	MaritalStatusID      string
	NumberOfDependents   uint8
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type CustomerAddresses struct {
	CorrespondenceAddress []Address
	ResidentialAddress    []Address
	WorkAddress           []Address
}

type Customer struct {
	Cid                     cid.Cid
	ContactInformation      CustomerContactInformation
	IdentificationDocuments CustomerIdentificationDocuments
	PersonalInformation     CustomerPersonalInformation
	Addresses               CustomerAddresses
	CreatedAt               time.Time
}
