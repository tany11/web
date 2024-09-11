package repository

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
	"log"

	"github.com/go-xorm/xorm"
)

type XormMasterRepository struct {
	engine *xorm.Engine
}

func NewXormMasterRepository(engine *xorm.Engine) repository.MasterRepository {
	return &XormMasterRepository{engine: engine}
}

func (r *XormMasterRepository) Create(master *entity.Master) error {
	_, err := r.engine.Insert(master)
	return err
}

func (r *XormMasterRepository) GetByID(id int64) (*entity.Master, error) {
	master := new(entity.Master)
	has, err := r.engine.ID(id).Get(master)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return master, nil
}

func (r *XormMasterRepository) GetByUsername(username string) (*entity.Master, error) {
	master := new(entity.Master)
	has, err := r.engine.Where("username = ?", username).Get(master)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return master, nil
}

func (r *XormMasterRepository) List(offset, limit int) ([]*entity.Master, error) {
	masters := make([]*entity.Master, 0)
	err := r.engine.Limit(limit, offset).Find(&masters)

	return masters, err
}

func (r *XormMasterRepository) Update(master *entity.Master) error {
	_, err := r.engine.ID(master.ID).Update(master)
	return err
}

func (r *XormMasterRepository) Delete(id int64) error {
	_, err := r.engine.ID(id).Delete(new(entity.Master))
	return err
}

func (r *XormMasterRepository) ListForDropdown(groupID int) ([]*entity.Master, error) {
	log.Println("ListForDropdown")
	masters := make([]*entity.Master, 0)
	err := r.engine.Cols("i_d", "master_name").Find(&masters)
	return masters, err
}

func (r *XormMasterRepository) ListForUsage(groupID int) ([]*entity.Master, error) {
	masters := make([]*entity.Master, 0)
	err := r.engine.Where("classification_name = ?", "UsageType").Find(&masters)
	return masters, err
}
