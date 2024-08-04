package entity

import (
	"time"
)

type Orders struct {
	ID                int64
	GroupID           int
	StoreID           int
	PhoneNumber       string
	CustomerName      string
	CustomerID        int64
	ModelName         string
	ActualModel       string
	CourseMin         string
	Price             int
	PostalCode        string
	Address           string
	DriverID          string
	ReservationFee    int
	TransportationFee int
	TravelCost        int
	Media             int
	Notes             string
	CardstaffID       string
	OrderStaffID      string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CompletionFlg     string
	IsDeleted         string
}
