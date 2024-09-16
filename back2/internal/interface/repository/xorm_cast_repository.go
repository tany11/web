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

func (r *XormCastRepository) GetByCastIID(CastID string) (*entity.Cast, error) {
	cast := new(entity.Cast)
	has, err := r.engine.Where("cast_i_d = ?", CastID).Get(cast)
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
	err := r.engine.Where("group_i_d = ?", groupID).Limit(limit, offset).Find(&casts)
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

func (r *XormCastRepository) ListForDropdown(groupID int) ([]*entity.Cast, error) {
	casts := make([]*entity.Cast, 0)
	err := r.engine.Where("group_i_d = ?", groupID).Cols("i_d", "cast_i_d", "cast_name", "working_flg").Find(&casts)
	return casts, err
}

func (r *XormCastRepository) UpdateWorkingFlg(id int64) error {

	_, err := r.engine.ID(id).Update(&entity.Cast{WorkingFlg: "1"})
	return err
}

func (r *XormCastRepository) ResetAllWorkingFlags() error {
	_, err := r.engine.Exec("UPDATE cast SET working_flg = '0'")
	return err
}
