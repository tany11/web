package repository

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"

	"github.com/go-xorm/xorm"
)

type XormCustomerRepository struct {
	engine *xorm.Engine
}

func NewXormCustomerOrderRepository(engine *xorm.Engine) repository.CustomerRepository {
	return &XormCustomerRepository{engine: engine}
}

func (r *XormCustomerRepository) Create(customer *entity.Customer) error {
	_, err := r.engine.Insert(customer)
	return err
}

func (r *XormCustomerRepository) GetByID(id int64) (*entity.Customer, error) {
	order := new(entity.Customer)
	has, err := r.engine.ID(id).Get(order)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return order, nil
}

func (r *XormCustomerRepository) List(groupID, offset, limit int) ([]*entity.Customer, error) {
	orders := make([]*entity.Customer, 0)
	err := r.engine.Where("group_id = ?", groupID).Limit(limit, offset).Find(&orders)
	return orders, err
}

func (r *XormCustomerRepository) ListByCustomerID(customerID int) ([]*entity.Customer, error) {
	orders := make([]*entity.Customer, 0)
	err := r.engine.Where("customer_id = ?", customerID).Find(&orders)
	return orders, err
}

func (r *XormCustomerRepository) Update(order *entity.Customer) error {
	_, err := r.engine.ID(order.ID).Update(order)
	return err
}

func (r *XormCustomerRepository) Delete(id int64) error {
	_, err := r.engine.ID(id).Delete(new(entity.Customer))
	return err
}

func (r *XormCustomerRepository) GetByPhoneNumber(phoneNumber string) (*entity.Customer, error) {
	order := new(entity.Customer)
	has, err := r.engine.Where("phone_number = ?", phoneNumber).Get(order)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return order, nil
}
