package entity

import (
	"time"
)

type Orders struct {
	ID                int64     //オーダーのID
	GroupID           int64     //グループID
	StoreID           int64     //店舗ID
	PhoneNumber       string    //顧客電話番号
	UsageType         string    //利用種別
	CustomerName      string    //顧客名
	CustomerID        int64     //顧客ID
	ModelName         string    //モデル名
	ActualModel       string    //実際のモデル
	CourseMin         int       //コース分
	ExtraCourse       int       //追加コース
	ExtraTime         int       //延長時間
	Price             int       //価格
	OptionPrice100    int       //オプション価格100
	OptionPrice50     int       //オプション価格50
	City              string    //市区町村
	Address           string    //住所
	DriverID          string    //運転手ID
	ReservationFee    int       //予約料
	TransportationFee int       //交通費
	TravelCost        int       //交通費
	Media             int       //メディア料
	Notes             string    //メモ
	CardstaffID       string    //カードスタッフID
	OrderStaffID      string    //オーダースタッフID
	ScheduledTime     time.Time //予定時間
	CreatedAt         time.Time //作成日
	UpdatedAt         time.Time //更新日
	DummyStoreFlg     string    //ダミー店舗フラグ
	CompletionFlg     string    //完了フラグ
	IsDeleted         string    //削除フラグ
}
