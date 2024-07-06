package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/errors"
	"back2/internal/domain/repository"
)

type OrderUseCase struct {
	orderRepo    repository.OrderRepository
	customerRepo repository.CustomerRepository
}

func NewOrderUseCase(orderRepo repository.OrderRepository, customerRepo repository.CustomerRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:    orderRepo,
		customerRepo: customerRepo,
	}
}

func (uc *OrderUseCase) Create(order *entity.Orders) error {
	customer, err := uc.customerRepo.GetByID(int64(order.CustomerID))
	if err != nil {
		return err
	}
	if customer == nil {
		return errors.ErrCustomerNotFound
	}

	return uc.orderRepo.Create(order)
}

func (uc *OrderUseCase) List(groupID, page, pageSize int) ([]*entity.Orders, error) {
	offset := (page - 1) * pageSize
	return uc.orderRepo.List(groupID, offset, pageSize)
}

func (uc *OrderUseCase) GetByID(id int64) (*entity.Orders, error) {
	return uc.orderRepo.GetByID(id)
}

// func (uc *OrderUseCase) List(groupID int, page, pageSize int) ([]*entity.Order, error) {
// 	offset := (page - 1) * pageSize
// 	return uc.orderRepo.List(groupID, offset, pageSize)
// }

func (uc *OrderUseCase) Update(order *entity.Orders) error {
	return uc.orderRepo.Update(order)
}

func (uc *OrderUseCase) Delete(id int64) error {
	return uc.orderRepo.Delete(id)
}

func (uc *OrderUseCase) GetByCustomerID(customerID int) ([]*entity.Orders, error) {
	return uc.orderRepo.ListByCustomerID(customerID)
}
