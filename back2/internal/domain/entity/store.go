package entity

import "time"

type Store struct {
	ID            int64     `xorm:"pk autoincr"`
	GroupID       int64     `xorm:"notnull"`
	StoreCode     int64     `xorm:"notnull"`
	StoreName     string    `xorm:"varchar(30) notnull"`
	DummyStoreFlg string    `xorm:"varchar(1) notnull default '0'"`
	CreatedAt     time.Time `xorm:"created"`
	UpdatedAt     time.Time `xorm:"updated"`
}
