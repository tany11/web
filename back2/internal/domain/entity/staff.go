package entity

import "time"

type Staff struct {
	ID             int64
	GroupID        int
	StaffID        string
	StaffLastName  string
	StaffFirstName string
	PasswordHash   string
	LineID         string
	OfficeFlg      string
	DriverFlg      string
	WebFlg         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	LastLogin      time.Time
	IsDeleted      string //削除フラグ
	Birthdate      time.Time
}
