package repository

import "back2/internal/domain/entity"

type StoreRepository interface {
	Create(store *entity.Store) error
	GetByID(id int64) (*entity.Store, error)
	List(offset, limit int) ([]*entity.Store, error)
	Update(store *entity.Store) error
	Delete(id int64) error
	ListForDropdown(groupID int) ([]*entity.Store, error)
}
