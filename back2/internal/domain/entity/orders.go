package entity

import (
	"time"
)

type Orders struct {
	ID             int64
	GroupID        int
	StoreID        int
	CustomerID     int
	CustomerName   string
	ModelName      string
	CastID         int
	OrderAmount    int
	PostalCode     string
	Address        string
	DriverID       int
	ReservationFee int
	TravelCost     int
	OtherFee       int
	OrderDate      time.Time
	UpdatedAt      time.Time
	IsDeleted      string
}
