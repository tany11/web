package repository

import "back2/internal/domain/entity"

type MasterRepository interface {
	Create(master *entity.Master) error
	GetByID(id int64) (*entity.Master, error)
	List(offset, limit int) ([]*entity.Master, error)
	Update(master *entity.Master) error
	Delete(id int64) error
	ListForDropdown(groupID int) ([]*entity.Master, error)
}
