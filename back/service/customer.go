package service

import (
	"back/models"
)

type CustomerService struct{}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) Create(customer *models.Customer) error {
	_, err := DbEngine.Insert(customer)
	return err
}

func (s *CustomerService) GetByID(id int64) (*models.Customer, error) {
	customer := &models.Customer{ID: id}
	has, err := DbEngine.Get(customer)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return customer, nil
}

func (s *CustomerService) GetByPhone(phonenumber int) (*models.Customer, error) {
	customer := &models.Customer{PhoneNumber: phonenumber}
	has, err := DbEngine.Get(customer)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return customer, nil
}

func (s *CustomerService) Update(customer *models.Customer) error {
	_, err := DbEngine.ID(customer.ID).Update(customer)
	return err
}

func (s *CustomerService) Delete(id int64) error {
	_, err := DbEngine.ID(id).Delete(&models.Customer{})
	return err
}

func (s *CustomerService) List(page, pageSize int) ([]*models.Customer, error) {
	customers := make([]*models.Customer, 0)
	err := DbEngine.Limit(pageSize, (page-1)*pageSize).Find(&customers)
	return customers, err
}
