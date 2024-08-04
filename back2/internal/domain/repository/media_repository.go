package repository

import "back2/internal/domain/entity"

type MediaRepository interface {
	Create(media *entity.Media) error
	GetByID(id int64) (*entity.Media, error)
	List(offset, limit int) ([]*entity.Media, error)
	Update(media *entity.Media) error
	Delete(id int64) error
	ListForDropdown(groupID int) ([]*entity.Media, error)
}
