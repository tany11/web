package repository

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
	"log"
	"time"

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
	// 24時間前の時刻を計算
	twentyFourHoursAgo := time.Now().Add(-24 * time.Hour)
	err := r.engine.Where("group_i_d = ?", groupID).
		And("completion_flg = 1").
		And("created_at > ?", twentyFourHoursAgo).
		Limit(limit, offset).
		Find(&orders)
	return orders, err
}
func (r *XormOrderRepository) ListReserved(groupID, offset, limit int) ([]*entity.Orders, error) {
	orders := make([]*entity.Orders, 0)
	err := r.engine.Where("group_i_d = ?", groupID).Where("completion_flg = 0").Where("is_deleted = 0").Limit(limit, offset).Find(&orders)
	return orders, err
}

func (r *XormOrderRepository) ListByCustomerID(customerID int) ([]*entity.Orders, error) {
	orders := make([]*entity.Orders, 0)
	err := r.engine.Where("customer_i_d = ?", customerID).Find(&orders)
	return orders, err
}

func (r *XormOrderRepository) Update(order *entity.Orders) error {
	session := r.engine.ID(order.ID)

	if order.ActualModel == "" {
		session = session.Cols("actual_model")
	}

	if !order.ScheduledTime.IsZero() {
		session = session.Cols("scheduled_time")
	}

	_, err := session.Update(order)
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

func (r *XormOrderRepository) GetByCustomerID(customerID int) (*entity.CustomerOrder, error) {
	customerOrder := new(entity.CustomerOrder)
	err := r.engine.Where("customer_i_d = ?", customerID).Find(&customerOrder)
	return customerOrder, err
}

func (r *XormOrderRepository) GetTotalPriceAndUseTime(customerID int) (int, int, error) {
	totalPrice, totalUseTime, err := r.getTotalPriceAndUseTimeImpl(customerID)
	return int(totalPrice), int(totalUseTime), err
}

func (r *XormOrderRepository) getTotalPriceAndUseTimeImpl(customerID int) (int64, int64, error) {
	totalPrice, err := r.engine.
		Where("customer_i_d = ?", customerID).
		SumInt(new(entity.Orders), "price")
	if err != nil {
		return 0, 0, err
	}

	ToralUseTime, err := r.engine.
		Where("customer_i_d = ?", customerID).
		Count(new(entity.Orders))

	return totalPrice, ToralUseTime, err
}

func (r *XormOrderRepository) ListSchedule(startDate, endDate string) ([]*entity.Orders, error) {

	orders := make([]*entity.Orders, 0)
	err := r.engine.Where("scheduled_time >= ?", startDate).
		And("scheduled_time <= ?", endDate).
		Find(&orders)
	return orders, err
}
