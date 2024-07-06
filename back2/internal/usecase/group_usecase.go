package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
)

type GroupUseCase struct {
	repo repository.GroupRepository
}

func NewGroupUseCase(repo repository.GroupRepository) *GroupUseCase {
	return &GroupUseCase{repo: repo}
}

func (uc *GroupUseCase) Create(group *entity.Group) error {
	return uc.repo.Create(group)
}

func (uc *GroupUseCase) GetByID(id int64) (*entity.Group, error) {
	return uc.repo.GetByID(id)
}

func (uc *GroupUseCase) GetByUsername(username string) (*entity.Group, error) {
	return uc.repo.GetByUsername(username)
}

func (uc *GroupUseCase) List(page, pageSize int) ([]*entity.Group, error) {
	offset := (page - 1) * pageSize
	return uc.repo.List(offset, pageSize)
}

func (uc *GroupUseCase) Update(group *entity.Group) error {
	return uc.repo.Update(group)
}

func (uc *GroupUseCase) Delete(id int64) error {
	return uc.repo.Delete(id)
}
