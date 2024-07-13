package entity

import "time"

type Cast struct {
	ID           int64     `xorm:"pk autoincr"`
	GroupID      int       `xorm:"notnull"`
	CastID       string    `xorm:"varchar(7) notnull unique"`
	CastName     string    `xorm:"varchar(255) notnull"`
	PasswordHash string    `xorm:"varchar(255) notnull"`
	LineID       string    `xorm:"varchar(255) unique notnull"`
	CreatedAt    time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
	LastLogin    time.Time `xorm:"timestamp"`
	Status       string    `xorm:"varchar(20) default 'active'"`
	Birthdate    time.Time `xorm:"date" json:"birthdate" time_format:"2006-01-02"`
}
