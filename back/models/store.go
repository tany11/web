package models

import (
	"time"
)

type Store struct {
	ID          int64     `xorm:"pk autoincr"`                  //id
	GroupID     int       `xorm:"notnull"`                      //グループID
	StoreName   string    `xorm:"varchar(255) notnull"`         //店舗名
	Address     string    `xorm:"text"`                         //住所
	PhoneNumber string    `xorm:"varchar(20)"`                  //電話番号
	Email       string    `xorm:"varchar(255) unique"`          //メールアドレス
	CreatedAt   time.Time `xorm:"created"`                      //登録日
	UpdatedAt   time.Time `xorm:"updated"`                      //更新日
	Status      string    `xorm:"varchar(20) default 'active'"` //ステータス
}
