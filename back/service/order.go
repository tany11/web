package service

import (
	"back/models"
)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) Create(order *models.Order) error {
	_, err := DbEngine.Insert(order)
	return err
}

func (s *OrderService) GetByID(id int64) (*models.Order, error) {
	order := &models.Order{ID: id}
	has, err := DbEngine.Get(order)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return order, nil
}

func (s *OrderService) Update(order *models.Order) error {
	_, err := DbEngine.ID(order.ID).Update(order)
	return err
}

func (s *OrderService) Delete(id int64) error {
	_, err := DbEngine.ID(id).Delete(&models.Order{})
	return err
}

func (s *OrderService) List(page, pageSize int) ([]*models.Order, error) {
	orders := make([]*models.Order, 0)
	err := DbEngine.Limit(pageSize, (page-1)*pageSize).Find(&orders)
	return orders, err
}
