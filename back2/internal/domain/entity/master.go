package entity

import "time"

type Master struct {
	ID         int64     `xorm:"pk autoincr"`
	StatusName string    `xorm:"varchar(255) notnull"`
	StatusCode string    `xorm:"pk autoincr"`
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}
