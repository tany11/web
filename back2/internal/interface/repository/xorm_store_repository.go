package repository

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
	"log"

	"github.com/go-xorm/xorm"
)

type XormStoreRepository struct {
	engine *xorm.Engine
}

func NewXormStoreRepository(engine *xorm.Engine) repository.StoreRepository {
	return &XormStoreRepository{engine: engine}
}

func (r *XormStoreRepository) Create(store *entity.Store) error {
	_, err := r.engine.Insert(store)
	return err
}

func (r *XormStoreRepository) GetByID(id int64) (*entity.Store, error) {
	store := new(entity.Store)
	has, err := r.engine.ID(id).Get(store)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return store, nil
}

func (r *XormStoreRepository) List(offset, limit int) ([]*entity.Store, error) {
	stores := make([]*entity.Store, 0)
	err := r.engine.Limit(limit, offset).Find(&stores)
	return stores, err
}

func (r *XormStoreRepository) Update(store *entity.Store) error {
	_, err := r.engine.ID(store.ID).Update(store)
	return err
}

func (r *XormStoreRepository) Delete(id int64) error {
	_, err := r.engine.ID(id).Delete(new(entity.Store))
	return err
}

func (r *XormStoreRepository) ListForDropdown(groupID int) ([]*entity.Store, error) {
	log.Println("ListForDropdown")
	stores := make([]*entity.Store, 0)
	err := r.engine.Where("group_i_d = ?", groupID).Cols("i_d", "store_name", "store_code").Find(&stores)
	log.Println(stores)
	return stores, err
}
