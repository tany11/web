package repository

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
	"log"

	"github.com/go-xorm/xorm"
)

type XormOrderRepository struct {
	engine *xorm.Engine
}

func NewXormOrderRepository(engine *xorm.Engine) repository.OrderRepository {
	return &XormOrderRepository{engine: engine}
}

func (r *XormOrderRepository) Create(order *entity.Orders) error {
	_, err := r.engine.Insert(order)
	if err != nil {
		log.Printf("データベースへの注文挿入に失敗しました。エラー: %v", err)
		return err
	}
	return nil
}

func (r *XormOrderRepository) GetByID(id int64) (*entity.Orders, error) {
	order := new(entity.Orders)
	has, err := r.engine.ID(id).Get(order)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return order, nil
}

func (r *XormOrderRepository) List(groupID, offset, limit int) ([]*entity.Orders, error) {
	orders := make([]*entity.Orders, 0)
	err := r.engine.Where("group_i_d = ?", groupID).Limit(limit, offset).Find(&orders)
	return orders, err
}

func (r *XormOrderRepository) ListReserved(groupID, offset, limit int) ([]*entity.Orders, error) {
	orders := make([]*entity.Orders, 0)
	err := r.engine.Where("group_i_d = ?", groupID).Where("completion_flg = 0").Where("is_deleted = 0").Limit(limit, offset).Find(&orders)
	return orders, err
}

func (r *XormOrderRepository) ListByCustomerID(customerID int) ([]*entity.Orders, error) {
	orders := make([]*entity.Orders, 0)
	err := r.engine.Where("customer_id = ?", customerID).Find(&orders)
	return orders, err
}

func (r *XormOrderRepository) Update(order *entity.Orders) error {
	_, err := r.engine.ID(order.ID).Update(order)
	return err
}

func (r *XormOrderRepository) Delete(id int64) error {
	_, err := r.engine.ID(id).Delete(new(entity.Orders))
	return err
}

func (r *XormOrderRepository) UpdateCompletionFlg(id int64) error {
	_, err := r.engine.ID(id).Update(new(entity.Orders))
	return err
}

func (r *XormOrderRepository) UpdateIsDeleted(id int64) error {
	_, err := r.engine.ID(id).Update(new(entity.Orders))
	return err
}
