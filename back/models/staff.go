package models

import (
	"time"
)

type StaffLogin struct {
	ID             int64     `xorm:"pk autoincr"`                  //id
	GroupID        int       `xorm:"notnull"`                      //グループID
	StaffID        string    `xorm:"varchar(7) notnull unique"`    //スタッフID
	StaffLastName  string    `xorm:"varchar(255) notnull"`         //スタッフ名(姓)
	StaffFirstName string    `xorm:"varchar(255) notnull"`         //スタッフ名(名)
	PasswordHash   string    `xorm:"varchar(255) notnull"`         //パスワードハッシュ
	Email          string    `xorm:"varchar(255) "`                //メールアドレス
	LineID         string    `xorm:"varchar(255) unique notnull"`  //ラインID
	PhoneNumber    string    `xorm:"varchar(20)"`                  //電話番号
	Position       string    `xorm:"varchar(255)"`                 //役割
	CreatedAt      time.Time `xorm:"created"`                      //登録日
	UpdatedAt      time.Time `xorm:"updated"`                      //更新日
	LastLogin      time.Time `xorm:"timestamp"`                    //最後のログイン日
	Status         string    `xorm:"varchar(20) default 'active'"` //ステータス
	Birthdate      time.Time `xorm:"date"`                         //誕生日
}
