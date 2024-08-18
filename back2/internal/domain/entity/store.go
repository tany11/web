package entity

import "time"

type Store struct {
	ID        int64     `xorm:"pk autoincr"`
	GroupID   int64     `xorm:"notnull"`
	StoreCode int64     `xorm:"notnull"`
	StoreName string    `xorm:"varchar(30) notnull"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
