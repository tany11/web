package entity

import "time"

type Customer struct {
	ID           int64     `xorm:"pk autoincr"`          //id
	GroupID      int       `xorm:"notnull"`              //グループID
	CustomerName string    `xorm:"varchar(255) notnull"` //顧客名
	PhoneNumber  string    `xorm:"varchar(255) notnull"` //電話番号
	Address      string    `xorm:"text"`                 //住所1
	Memo         string    `xorm:"text"`                 //メモ
	CreatedAt    time.Time `xorm:"created"`              //登録日
	UpdatedAt    time.Time `xorm:"updated"`              //更新日
}
