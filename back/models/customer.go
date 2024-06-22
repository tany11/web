package models

import (
	"time"
)

type Customer struct {
	ID           int64     `xorm:"pk autoincr"`          //id
	GroupID      int       `xorm:"notnull"`              //グループID
	CustomerName string    `xorm:"varchar(255) notnull"` //顧客名
	PhoneNumber  string    `xorm:"varchar(20)"`          //電話番号
	Email        string    `xorm:"varchar(255) unique"`  //メールアドレス
	Address      string    `xorm:"text"`                 //住所
	CreatedAt    time.Time `xorm:"created"`              //登録日
	UpdatedAt    time.Time `xorm:"updated"`              //更新日
}
