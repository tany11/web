package entity

import "time"

type Master struct {
	ID         int64  `xorm:"pk autoincr"`
	StatusName string `xorm:"varchar(30) notnull"`
	StatusCode string
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}
