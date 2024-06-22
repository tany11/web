package models

import (
	"time"
)

type GroupLogin struct {
	ID           int64     `xorm:"pk autoincr"`                  //id
	GroupName    string    `xorm:"varchar(255) notnull"`         //グループ名
	Username     string    `xorm:"varchar(255) unique notnull"`  //ユーザー名
	PasswordHash string    `xorm:"varchar(255) notnull"`         //パスワードハッシュ
	Email        string    `xorm:"varchar(255) unique notnull"`  //メールアドレス
	PhoneNumber  string    `xorm:"varchar(20)"`                  //電話番号
	Address      string    `xorm:"text"`                         //住所
	CreatedAt    time.Time `xorm:"created"`                      //登録日
	UpdatedAt    time.Time `xorm:"updated"`                      //更新日
	LastLogin    time.Time `xorm:"timestamp"`                    //最後のログイン日
	Status       string    `xorm:"varchar(20) default 'active'"` //ステータス
}
