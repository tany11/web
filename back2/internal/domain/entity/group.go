package entity

import "time"

type Group struct {
	ID           int64
	GroupName    string
	Username     string
	PasswordHash string
	Email        string
	PhoneNumber  string
	Address      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastLogin    time.Time
	Status       string
}
