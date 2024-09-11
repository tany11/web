package repository

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
	"fmt"
	"log"
	"time"

	"github.com/go-xorm/xorm"
)

type XormTipsRepository struct {
	engine *xorm.Engine
}

func NewXormTipsRepository(engine *xorm.Engine) repository.TipsRepository {
	return &XormTipsRepository{engine: engine}
}

func (r *XormTipsRepository) Create(tips *entity.Tips) error {
	tips.IsDeleted = "0"
	tips.CompletionFlg = "0"
	_, err := r.engine.Insert(tips)
	if err != nil {
		log.Printf("データベースへのヒント挿入に失敗しました。エラー: %v", err)
		return err
	}
	return nil
}

func (r *XormTipsRepository) GetByID(id int64) (*entity.Tips, error) {
	tips := new(entity.Tips)
	has, err := r.engine.ID(id).Get(tips)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return tips, nil
}

func (r *XormTipsRepository) List(groupID, offset, limit int) ([]*entity.Tips, error) {
	tips := make([]*entity.Tips, 0)
	// 24時間前の時刻を計算
	twentyFourHoursAgo := time.Now().Add(-24 * time.Hour)
	err := r.engine.Where("group_i_d = ?", groupID).
		And("created_at > ?", twentyFourHoursAgo).
		Limit(limit, offset).
		Find(&tips)
	return tips, err
}

func (r *XormTipsRepository) ListReserved(groupID, offset, limit int) ([]*entity.Tips, error) {
	tips := make([]*entity.Tips, 0)
	err := r.engine.Where("group_i_d = ?", groupID).Where("completion_flg = 0").Where("is_deleted = 0").Limit(limit, offset).Find(&tips)
	return tips, err
}

func (r *XormTipsRepository) Update(tips *entity.Tips) error {
	session := r.engine.ID(tips.ID)
	fmt.Println("ちっぷす", tips)

	session = session.Cols("actual_model")

	if !tips.ScheduledTime.IsZero() {
		session = session.Cols("scheduled_time")
	}

	if tips.ScheduledBox != "" {
		session = session.Cols("scheduled_box")
	}

	if tips.Title != "" {
		session = session.Cols("title")
	}

	if tips.Notes != "" {
		session = session.Cols("notes")
	}

	if tips.CompletionFlg != "" {
		session = session.Cols("completion_flg")
	}

	if tips.IsDeleted != "" {
		session = session.Cols("is_deleted")
	}

	_, err := session.Update(tips)
	return err
}

func (r *XormTipsRepository) Delete(id int64) error {
	_, err := r.engine.ID(id).Delete(new(entity.Tips))
	return err
}

func (r *XormTipsRepository) UpdateCompletionFlg(id int64, tips *entity.Tips) error {
	session := r.engine.ID(id)
	if tips.CompletionFlg != "" {
		session = session.Cols("completion_flg")
	}
	_, err := session.Update(tips)
	return err
}

func (r *XormTipsRepository) UpdateIsDeleted(id int64, tips *entity.Tips) error {
	session := r.engine.ID(id)
	if tips.IsDeleted != "" {
		session = session.Cols("is_deleted")
	}
	_, err := session.Update(tips)
	return err
}

func (r *XormTipsRepository) ListSchedule(groupID int, startDate, endDate string) ([]*entity.Tips, error) {

	tips := make([]*entity.Tips, 0)
	err := r.engine.Where("scheduled_time >= ?", startDate).
		And("scheduled_time <= ?", endDate).
		And("is_deleted = 0").
		Find(&tips)
	return tips, err
}

func (r *XormTipsRepository) ListByCustomerID(customerID int64) ([]*entity.Tips, error) {
	tips := make([]*entity.Tips, 0)
	err := r.engine.Where("customer_id = ?", customerID).Find(&tips)
	return tips, err
}
