package entity

import "time"

type Cast struct {
	ID            int64     `xorm:"pk autoincr"`
	GroupID       int       `xorm:"notnull"`
	CastID        string    `xorm:"varchar(7) notnull unique"`
	LastName      string    `xorm:"varchar(30) notnull"`
	FirstName     string    `xorm:"varchar(30) notnull"`
	CastName      string    `xorm:"varchar(30) notnull"`
	PasswordHash  string    `xorm:"varchar(255) notnull"`
	LineID        string    `xorm:"varchar(30) unique notnull"`
	EffectiveDate time.Time `xorm:"date"`
	CreatedAt     time.Time `xorm:"created"`
	UpdatedAt     time.Time `xorm:"updated"`
	LastLogin     time.Time `xorm:"timestamp"`
	IsDeleted     string    `xorm:"varchar(1) notnull default('0')"` //削除フラグ
	Birthdate     time.Time `xorm:"date" json:"birthdate" time_format:"2006-01-02"`
}
