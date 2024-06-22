package service

import (
	"back/models"
)

type GroupLoginService struct{}

func NewGroupLoginService() *GroupLoginService {
	return &GroupLoginService{}
}

func (s *GroupLoginService) Create(groupLogin *models.GroupLogin) error {
	_, err := DbEngine.Insert(groupLogin)
	if err != nil {
		return err
	}
	return nil
}

func (s *GroupLoginService) GetByID(id int64) (*models.GroupLogin, error) {
	groupLogin := &models.GroupLogin{ID: id}
	has, err := DbEngine.Get(groupLogin)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return groupLogin, nil
}

func (s *GroupLoginService) Update(groupLogin *models.GroupLogin) error {
	_, err := DbEngine.ID(groupLogin.ID).Update(groupLogin)
	return err
}

func (s *GroupLoginService) Delete(id int64) error {
	_, err := DbEngine.ID(id).Delete(&models.GroupLogin{})
	return err
}

func (s *GroupLoginService) List(page, pageSize int) ([]*models.GroupLogin, error) {
	groupLogins := make([]*models.GroupLogin, 0)
	err := DbEngine.Limit(pageSize, (page-1)*pageSize).Find(&groupLogins)
	return groupLogins, err
}
