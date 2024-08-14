package repository

import "back2/internal/domain/entity"

type CustomerRepository interface {
	Create(order *entity.Customer) error
	GetByID(id int64) (*entity.Customer, error)
	List(groupID, offset, limit int) ([]*entity.Customer, error)
	ListByCustomerID(customerID int) ([]*entity.Customer, error)
	Update(order *entity.Customer) error
	Delete(id int64) error
	GetByPhoneNumber(phoneNumber string) (*entity.Customer, error)
	GetSearchList(customerSearch entity.CustomerSearch) ([]*entity.CustomerOrder, error)
}
