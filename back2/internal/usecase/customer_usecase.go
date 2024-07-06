package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
)

type CustomerUseCase struct {
	repo repository.CustomerRepository
}

func NewCustomerUseCase(repo repository.CustomerRepository) *CustomerUseCase {
	return &CustomerUseCase{repo: repo}
}

func (uc *CustomerUseCase) Create(customer *entity.Customer) error {
	return uc.repo.Create(customer)
}

func (uc *CustomerUseCase) GetByID(id int64) (*entity.Customer, error) {
	return uc.repo.GetByID(id)
}

func (uc *CustomerUseCase) GetByPhoneNumber(phoneNumber string) (*entity.Customer, error) {
	return uc.repo.GetByPhoneNumber(phoneNumber)
}

func (uc *CustomerUseCase) List(groupID int, page, pageSize int) ([]*entity.Customer, error) {
	offset := (page - 1) * pageSize
	return uc.repo.List(groupID, offset, pageSize)
}

func (uc *CustomerUseCase) Update(customer *entity.Customer) error {
	return uc.repo.Update(customer)
}

func (uc *CustomerUseCase) Delete(id int64) error {
	return uc.repo.Delete(id)
}
