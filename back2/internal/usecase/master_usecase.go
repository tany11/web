package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
)

type MasterUseCase struct {
	repo repository.MasterRepository
}

func NewMasterUseCase(repo repository.MasterRepository) *MasterUseCase {
	return &MasterUseCase{repo: repo}
}

func (uc *MasterUseCase) Create(master *entity.Master) error {
	return uc.repo.Create(master)
}

func (uc *MasterUseCase) GetByID(id int64) (*entity.Master, error) {
	return uc.repo.GetByID(id)
}

func (uc *MasterUseCase) List(page, pageSize int) ([]*entity.Master, error) {
	offset := (page - 1) * pageSize
	return uc.repo.List(offset, pageSize)
}

func (uc *MasterUseCase) Update(master *entity.Master) error {
	return uc.repo.Update(master)
}

func (uc *MasterUseCase) Delete(id int64) error {
	return uc.repo.Delete(id)
}

func (uc *MasterUseCase) ListForDropdown(groupID int) ([]*entity.Master, error) {
	return uc.repo.ListForDropdown(groupID)
}
