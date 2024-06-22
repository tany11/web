package service

import (
	"back/models"
)

type StoreService struct{}

func NewStoreService() *StoreService {
	return &StoreService{}
}

func (s *StoreService) Create(store *models.Store) error {
	_, err := DbEngine.Insert(store)
	return err
}

func (s *StoreService) GetByID(id int64) (*models.Store, error) {
	store := &models.Store{ID: id}
	has, err := DbEngine.Get(store)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return store, nil
}

func (s *StoreService) Update(store *models.Store) error {
	_, err := DbEngine.ID(store.ID).Update(store)
	return err
}

func (s *StoreService) Delete(id int64) error {
	_, err := DbEngine.ID(id).Delete(&models.Store{})
	return err
}

func (s *StoreService) List(page, pageSize int) ([]*models.Store, error) {
	stores := make([]*models.Store, 0)
	err := DbEngine.Limit(pageSize, (page-1)*pageSize).Find(&stores)
	return stores, err
}
