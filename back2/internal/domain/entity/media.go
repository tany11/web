package entity

import "time"

type Media struct {
	ID        int64     `xorm:"pk autoincr"`
	MediaName string    `xorm:"varchar(255) notnull"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
