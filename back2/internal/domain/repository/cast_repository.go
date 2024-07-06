package repository

import "back2/internal/domain/entity"

type CastRepository interface {
	Create(cast *entity.Cast) error
	GetByID(id int64) (*entity.Cast, error)
	List(groupID, offset, limit int) ([]*entity.Cast, error)
	Update(cast *entity.Cast) error
	Delete(id int64) error
}
