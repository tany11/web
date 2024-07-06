package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
)

type CastUseCase struct {
	repo repository.CastRepository
}

func NewCastUseCase(repo repository.CastRepository) *CastUseCase {
	return &CastUseCase{repo: repo}
}

func (uc *CastUseCase) Create(cast *entity.Cast) error {
	return uc.repo.Create(cast)
}

func (uc *CastUseCase) GetByID(id int64) (*entity.Cast, error) {
	return uc.repo.GetByID(id)
}

func (uc *CastUseCase) List(groupID int, page, pageSize int) ([]*entity.Cast, error) {
	offset := (page - 1) * pageSize
	return uc.repo.List(groupID, offset, pageSize)
}

func (uc *CastUseCase) Update(cast *entity.Cast) error {
	return uc.repo.Update(cast)
}

func (uc *CastUseCase) Delete(id int64) error {
	return uc.repo.Delete(id)
}
