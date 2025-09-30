package customers

import "time"

type PhoneNumber struct {
	Id          string
	CountryCode string
	Number      string
	IsPrimary   bool
	Verified    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	VerifiedAt  time.Time
}
