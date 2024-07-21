package entity

import (
	"time"
)

type Orders struct {
	ID                int64
	GroupID           int `json:"group_id" xorm:"'group_i_d'"` // この行を確認//グループID
	StoreID           int
	PhoneNumber       string
	CustomerName      string
	CustomerID        int64
	ModelName         string
	ActualModel       string
	CourceMin         string
	Price             int
	PostalCode        string
	Address           string
	DriverID          string
	ReservationFee    int
	TransportationFee int
	TravelCost        int
	Media             string
	Notes             string
	CardstaffID       string
	OrderStaffID      string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CompletionFlg     string `default:"0"`
	IsDeleted         string `default:"0"`
}
