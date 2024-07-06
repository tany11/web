package repository

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"

	"github.com/go-xorm/xorm"
)

type XormCastRepository struct {
	engine *xorm.Engine
}

func NewXormCastRepository(engine *xorm.Engine) repository.CastRepository {
	return &XormCastRepository{engine: engine}
}

func (r *XormCastRepository) Create(cast *entity.Cast) error {
	_, err := r.engine.Insert(cast)
	return err
}

func (r *XormCastRepository) GetByID(id int64) (*entity.Cast, error) {
	cast := &entity.Cast{}
	has, err := r.engine.ID(id).Get(cast)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return cast, nil
}

func (r *XormCastRepository) List(groupID int, offset, limit int) ([]*entity.Cast, error) {
	casts := make([]*entity.Cast, 0)
	err := r.engine.Where("group_id = ?", groupID).Limit(limit, offset).Find(&casts)
	return casts, err
}

func (r *XormCastRepository) Update(cast *entity.Cast) error {
	_, err := r.engine.ID(cast.ID).Update(cast)
	return err
}

func (r *XormCastRepository) Delete(id int64) error {
	_, err := r.engine.ID(id).Delete(&entity.Cast{})
	return err
}
