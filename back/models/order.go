package models

import (
	"time"
)

type Order struct {
	ID          int64     `xorm:"pk autoincr"`           //id
	GroupID     int       `xorm:"notnull"`               //グループID
	CastID      int       `xorm:"notnull"`               //キャストID
	CustomerID  int       `xorm:"notnull"`               //顧客ID
	StoreID     int       `xorm:"notnull"`               //店舗ID
	OrderDate   time.Time `xorm:"created"`               //注文日
	OrderAmount float64   `xorm:"decimal(10,2) notnull"` //注文金額
}
