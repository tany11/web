package repository

import "back2/internal/domain/entity"

type OrderRepository interface {
	Create(order *entity.Orders) error
	GetByID(id int64) (*entity.Orders, error)
	List(groupID, offset, limit int) ([]*entity.Orders, error)
	ListByCustomerID(customerID int) ([]*entity.Orders, error)
	Update(order *entity.Orders) error
	Delete(id int64) error
	UpdateCompletionFlg(id int64) error
	ListReserved(groupID, offset, limit int) ([]*entity.Orders, error)
	UpdateIsDeleted(id int64) error
	GetTotalPriceAndUseTime(customerID int) (int, int, error)
	ListSchedule(groupID int, startDate, endDate string) ([]*entity.Orders, error)
	UpdateSchedule(order *entity.Orders) error
	CalcPayroll(id int64) error
}
