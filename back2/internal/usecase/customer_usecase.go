package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
	"log"
)

type CustomerUseCase struct {
	repo      repository.CustomerRepository
	repoOrder repository.OrderRepository
}

func NewCustomerUseCase(repo repository.CustomerRepository, repoOrder repository.OrderRepository) *CustomerUseCase {
	return &CustomerUseCase{repo: repo, repoOrder: repoOrder}
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

func (uc *CustomerUseCase) GetSearchList(customerSearch entity.CustomerSearch) ([]*entity.CustomerOrder, error) {
	log.Printf("UseCase: 顧客検索開始 - パラメータ: %+v", customerSearch)
	customers, err := uc.repo.GetSearchList(customerSearch)
	if err != nil {
		log.Printf("UseCase: 顧客検索エラー: %v", err)
		return nil, err
	}
	log.Printf("UseCase: 顧客検索完了 - 結果件数: %d", len(customers))
	return customers, nil
}
