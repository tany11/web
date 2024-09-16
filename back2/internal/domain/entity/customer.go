package entity

import (
	"time"
)

type Customer struct {
	ID           int64     `xorm:"pk autoincr"`                     //id
	GroupID      int64     `xorm:"notnull"`                         //グループID
	CustomerName string    `xorm:"varchar(30) notnull"`             //顧客名
	PhoneNumber  string    `xorm:"varchar(15) notnull"`             //電話番号
	City1        string    `xorm:"text"`                            //市区町村1
	Address1     string    `xorm:"text"`                            //住所1
	City2        string    `xorm:"text"`                            //市区町村2
	Address2     string    `xorm:"text"`                            //住所2
	City3        string    `xorm:"text"`                            //市区町村3
	Address3     string    `xorm:"text"`                            //住所3
	Memo         string    `xorm:"text"`                            //メモ
	CreatedAt    time.Time `xorm:"created"`                         //登録日
	UpdatedAt    time.Time `xorm:"updated"`                         //更新日
	IsDeleted    string    `xorm:"varchar(1) notnull default('0')"` //削除フラグ
}

type CustomerOrder struct {
	ID           int64            `xorm:"pk autoincr"`         //id
	GroupID      int64            `xorm:"notnull"`             //グループID
	CustomerName string           `xorm:"varchar(30) notnull"` //顧客名
	PhoneNumber  string           `xorm:"varchar(15) notnull"` //電話番号
	City1        string           `xorm:"text"`                //市区町村1
	Address1     string           `xorm:"text"`                //住所1
	City2        string           `xorm:"text"`                //市区町村2
	Address2     string           `xorm:"text"`                //住所2
	City3        string           `xorm:"text"`                //市区町村3
	Address3     string           `xorm:"text"`                //住所3
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

type CustomerSearch struct {
	PhoneLast4  string
	CastID      string
	StoreID     string
	CreatedFrom string
	CreatedTo   string
}
