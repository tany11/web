package models

import (
	"time"
)

type Order struct {
	ID             int64     `xorm:"pk autoincr"`                  //id
	GroupID        int       `xorm:"notnull"`                      //グループID
	StoreID        int       `xorm:"notnull"`                      //店舗ID
	CustomerID     int       `xorm:"notnull"`                      //顧客ID
	CustomerName   string    `xorm:"notnull"`                      //顧客名
	ModelName      string    `xorm:"text"`                         //モデル
	CastID         int       `xorm:"notnull"`                      //キャストID
	OrderAmount    int       `xorm:"notnull"`                      //注文金額
	PostalCode     string    `xorm:"text"`                         //郵便番号
	Address        string    `xorm:"text"`                         //住所
	DriverID       int       `xorm:"notnull"`                      //ドライバー
	ReservationFee int       `xorm:"notnull"`                      //指名料
	TravelCost     int       `xorm:"notnull"`                      //交通費
	OtherFee       int       `xorm:"notnull"`                      //出張費
	OrderDate      time.Time `xorm:"created"`                      //注文日
	UpdatedAt      time.Time `xorm:"updated"`                      //更新日
	Status         string    `xorm:"varchar(20) default 'active'"` //ステータス

}
