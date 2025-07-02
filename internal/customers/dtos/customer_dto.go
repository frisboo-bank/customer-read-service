package dtos

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type CustomerDto struct {
	Id                  uuid.UUID
	FirstName           string
	MiddleName          string
	LastName            string
	FirstNameInEnglish  string
	MiddleNameInEnglish string
	LastNameInEnglish   string
	DateOfBirth         string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	OriginalVersion     int64
}
