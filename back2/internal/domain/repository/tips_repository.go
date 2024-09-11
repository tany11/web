package repository

import (
	"back2/internal/domain/entity"
)

type TipsRepository interface {
	Create(tips *entity.Tips) error
	GetByID(id int64) (*entity.Tips, error)
	List(groupID, offset, limit int) ([]*entity.Tips, error)
	Update(tips *entity.Tips) error
	Delete(id int64) error
	UpdateCompletionFlg(id int64, tips *entity.Tips) error
	UpdateIsDeleted(id int64, tips *entity.Tips) error
	ListReserved(groupID, offset, limit int) ([]*entity.Tips, error)
	ListSchedule(groupID int, startDate, endDate string) ([]*entity.Tips, error)
}
