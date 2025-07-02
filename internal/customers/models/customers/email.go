package models

import "time"

type Email struct {
	Id         string
	Email      string
	IsPrimary  bool
	Verified   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	VerifiedAt time.Time
}
