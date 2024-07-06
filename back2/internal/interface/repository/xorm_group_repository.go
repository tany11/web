package repository

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"

	"github.com/go-xorm/xorm"
)

type XormGroupRepository struct {
	engine *xorm.Engine
}

func NewXormGroupRepository(engine *xorm.Engine) repository.GroupRepository {
	return &XormGroupRepository{engine: engine}
}

func (r *XormGroupRepository) Create(group *entity.Group) error {
	_, err := r.engine.Insert(group)
	return err
}

func (r *XormGroupRepository) GetByID(id int64) (*entity.Group, error) {
	group := new(entity.Group)
	has, err := r.engine.ID(id).Get(group)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return group, nil
}

func (r *XormGroupRepository) GetByUsername(username string) (*entity.Group, error) {
	group := new(entity.Group)
	has, err := r.engine.Where("username = ?", username).Get(group)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return group, nil
}

func (r *XormGroupRepository) List(offset, limit int) ([]*entity.Group, error) {
	groups := make([]*entity.Group, 0)
	err := r.engine.Limit(limit, offset).Find(&groups)
	return groups, err
}

func (r *XormGroupRepository) Update(group *entity.Group) error {
	_, err := r.engine.ID(group.ID).Update(group)
	return err
}

func (r *XormGroupRepository) Delete(id int64) error {
	_, err := r.engine.ID(id).Delete(new(entity.Group))
	return err
}
