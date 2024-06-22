package service

import (
	"back/models"
)

type CastLoginService struct{}

func NewCastLoginService() *CastLoginService {
	return &CastLoginService{}
}

func (s *CastLoginService) Create(castLogin *models.CastLogin) error {
	_, err := DbEngine.Insert(castLogin)
	return err
}

func (s *CastLoginService) GetByID(id int64) (*models.CastLogin, error) {
	castLogin := &models.CastLogin{ID: id}
	has, err := DbEngine.Get(castLogin)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return castLogin, nil
}

func (s *CastLoginService) Update(castLogin *models.CastLogin) error {
	_, err := DbEngine.ID(castLogin.ID).Update(castLogin)
	return err
}

func (s *CastLoginService) Delete(id int64) error {
	_, err := DbEngine.ID(id).Delete(&models.CastLogin{})
	return err
}

func (s *CastLoginService) List(page, pageSize int) ([]*models.CastLogin, error) {
	castLogins := make([]*models.CastLogin, 0)
	err := DbEngine.Limit(pageSize, (page-1)*pageSize).Find(&castLogins)
	return castLogins, err
}
