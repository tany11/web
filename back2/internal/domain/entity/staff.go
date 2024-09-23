package entity

import "time"

type Staff struct {
	ID             int64     `xorm:"pk autoincr 'i_d'" json:"id"`
	StaffID        string    `xorm:"varchar(7) notnull unique"`
	GroupID        int64     `xorm:"notnull"`
	StaffLastName  string    `xorm:"varchar(30) notnull"`
	StaffFirstName string    `xorm:"varchar(30) notnull"`
	PasswordHash   string    `xorm:"varchar(255) notnull"`
	LineID         string    `xorm:"varchar(30) unique notnull"`
	OfficeFlg      string    `xorm:"varchar(1) notnull default('0')"`
	DriverFlg      string    `xorm:"varchar(1) notnull default('0')"`
	WebFlg         string    `xorm:"varchar(1) notnull default('0')"`
	CreatedAt      time.Time `xorm:"created"`
	UpdatedAt      time.Time `xorm:"updated"`
	LastLogin      time.Time `xorm:"timestamp"`
	IsDeleted      string    `xorm:"varchar(1) notnull default('0')"` //削除フラグ
	Birthdate      time.Time `xorm:"date" json:"birthdate" time_format:"2006-01-02"`
}
