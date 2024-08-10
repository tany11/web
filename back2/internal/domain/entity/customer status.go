package entity

import "time"

type CustomerStatus struct {
	ID              int64     `xorm:"pk autoincr"`          //id
	CustomerID      int64     `xorm:"notnull"`              //顧客ID
	CustomerStatus1 string    `xorm:"varchar(255) notnull"` //顧客ステータス1
	CustomerStatus2 string    `xorm:"varchar(255) notnull"` //顧客ステータス2
	CustomerStatus3 string    `xorm:"varchar(255) notnull"` //顧客ステータス3
	CustomerStatus4 string    `xorm:"varchar(255) notnull"` //顧客ステータス4
	CustomerStatus5 string    `xorm:"varchar(255) notnull"` //顧客ステータス5
	CreatedAt       time.Time `xorm:"created"`              //登録日
	UpdatedAt       time.Time `xorm:"updated"`              //更新日
}
