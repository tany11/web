package repository

import "back2/internal/domain/entity"

type GroupRepository interface {
	Create(group *entity.Group) error
	GetByID(id int64) (*entity.Group, error)
	GetByUsername(username string) (*entity.Group, error)
	List(offset, limit int) ([]*entity.Group, error)
	Update(group *entity.Group) error
	Delete(id int64) error
}
