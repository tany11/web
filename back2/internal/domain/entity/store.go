package entity

import "time"

type Store struct {
	ID        int64     `xorm:"pk autoincr"`
	Name      string    `xorm:"varchar(255) notnull"`
	Address   string    `xorm:"varchar(255)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
