package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
)

type StoreUseCase struct {
	repo repository.StoreRepository
}

func NewStoreUseCase(repo repository.StoreRepository) *StoreUseCase {
	return &StoreUseCase{repo: repo}
}

func (uc *StoreUseCase) Create(store *entity.Store) error {
	return uc.repo.Create(store)
}

func (uc *StoreUseCase) GetByID(id int64) (*entity.Store, error) {
	return uc.repo.GetByID(id)
}

func (uc *StoreUseCase) List(page, pageSize int) ([]*entity.Store, error) {
	offset := (page - 1) * pageSize
	return uc.repo.List(offset, pageSize)
}

func (uc *StoreUseCase) Update(store *entity.Store) error {
	return uc.repo.Update(store)
}

func (uc *StoreUseCase) Delete(id int64) error {
	return uc.repo.Delete(id)
}

func (uc *StoreUseCase) ListForDropdown(groupID int) ([]*entity.Store, error) {
	return uc.repo.ListForDropdown(groupID)
}
