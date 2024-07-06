package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

var DBEngine *xorm.Engine

type CastLogin struct {
	ID           int64     `xorm:"pk autoincr"`                  // 自動インクリメントの整数ID
	GroupID      int       `xorm:"notnull"`                      //グループID
	CastID       string    `xorm:"varchar(7) notnull unique"`    //キャストID
	CastName     string    `xorm:"varchar(255) notnull"`         //キャスト名
	PasswordHash string    `xorm:"varchar(255) notnull"`         //パスワードハッシュ
	LineID       string    `xorm:"varchar(255) unique notnull"`  //ラインID
	CreatedAt    time.Time `xorm:"created"`                      //登録日
	UpdatedAt    time.Time `xorm:"updated"`                      //更新日
	LastLogin    time.Time `xorm:"timestamp"`                    //最後のログイン日
	Status       string    `xorm:"varchar(20) default 'active'"` //ステータス
	Birthdate    time.Time `xorm:"date"`                         //誕生日
}
