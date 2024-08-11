package entity

import (
	"time"
)

type Customer struct {
	ID           int64     `xorm:"pk autoincr"`         //id
	GroupID      int       `xorm:"notnull"`             //グループID
	CustomerName string    `xorm:"varchar(30) notnull"` //顧客名
	PhoneNumber  string    `xorm:"varchar(15) notnull"` //電話番号
	Address      string    `xorm:"text"`                //住所1
	Memo         string    `xorm:"text"`                //メモ
	CreatedAt    time.Time `xorm:"created"`             //登録日
	UpdatedAt    time.Time `xorm:"updated"`             //更新日
}

type CustomerOrder struct {
	ID           int64            `xorm:"pk autoincr"`         //id
	GroupID      int              `xorm:"notnull"`             //グループID
	CustomerName string           `xorm:"varchar(30) notnull"` //顧客名
	PhoneNumber  string           `xorm:"varchar(15) notnull"` //電話番号
	Address      string           `xorm:"text"`                //住所1
	Memo         string           `xorm:"text"`                //メモ
	TotalPrice   int              `xorm:"int"`                 //合計金額
	TotalUseTime int              `xorm:"int"`                 //合計時間
	OrderList    []Orders         `xorm:"-"`
	StatusList   []CustomerStatus `xorm:"-"`
	CreatedAt    time.Time        `xorm:"created"` //登録日
	UpdatedAt    time.Time        `xorm:"updated"` //更新日
}

type CustomerStatus struct {
	ID             int64     `xorm:"pk autoincr"`         //id
	CustomerID     int64     `xorm:"notnull"`             //顧客ID
	CustomerStatus string    `xorm:"varchar(25) notnull"` //顧客ステータス
	CreatedAt      time.Time `xorm:"created"`             //登録日
	UpdatedAt      time.Time `xorm:"updated"`             //更新日
}
