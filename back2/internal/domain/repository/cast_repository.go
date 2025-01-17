package repository

import "back2/internal/domain/entity"

type CastRepository interface {
	Create(cast *entity.Cast) error
	GetByID(id int64) (*entity.Cast, error)
	GetByCastIID(CastID string) (*entity.Cast, error)
	List(groupID int) ([]*entity.Cast, error)
	Update(cast *entity.Cast) error
	Delete(id int64) error
	ListForDropdown(groupID int) ([]*entity.Cast, error) // 新しいメソッドを追加
	UpdateWorkingFlg(id int64) error
	ResetAllWorkingFlags() error // 新しいメソッドを追加
}
