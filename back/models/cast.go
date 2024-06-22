package models

import (
	"time"
)

type CastLogin struct {
	ID             int64     `xorm:"pk autoincr"`                  //id
	GroupID        int       `xorm:"notnull"`                      //グループID
	CastName       string    `xorm:"varchar(255) notnull"`         //キャスト名
	Username       string    `xorm:"varchar(255) unique notnull"`  //ユーザー名
	PasswordHash   string    `xorm:"varchar(255) notnull"`         //パスワードハッシュ
	Email          string    `xorm:"varchar(255) unique notnull"`  //メールアドレス
	PhoneNumber    string    `xorm:"varchar(20)"`                  //電話番号
	Position       string    `xorm:"varchar(255)"`                 //役割
	Address        string    `xorm:"text"`                         //住所
	CreatedAt      time.Time `xorm:"created"`                      //登録日
	UpdatedAt      time.Time `xorm:"updated"`                      //更新日
	LastLogin      time.Time `xorm:"timestamp"`                    //最後のログイン日
	Status         string    `xorm:"varchar(20) default 'active'"` //ステータス
	ProfilePicture string    `xorm:"text"`                         //プロフィール画像
	Birthdate      time.Time `xorm:"date"`                         //誕生日
}
