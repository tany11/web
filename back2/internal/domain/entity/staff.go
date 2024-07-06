package entity

import "time"

type Staff struct {
	ID             int64
	GroupID        int
	StaffID        string
	StaffLastName  string
	StaffFirstName string
	PasswordHash   string
	Email          string
	LineID         string
	PhoneNumber    string
	Position       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	LastLogin      time.Time
	Status         string
	Birthdate      time.Time
}
