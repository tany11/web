package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
)

type MediaUseCase struct {
	repo repository.MediaRepository
}

func NewMediaUseCase(repo repository.MediaRepository) *MediaUseCase {
	return &MediaUseCase{repo: repo}
}

func (uc *MediaUseCase) Create(media *entity.Media) error {
	return uc.repo.Create(media)
}

func (uc *MediaUseCase) GetByID(id int64) (*entity.Media, error) {
	return uc.repo.GetByID(id)
}

func (uc *MediaUseCase) List(page, pageSize int) ([]*entity.Media, error) {
	offset := (page - 1) * pageSize
	return uc.repo.List(offset, pageSize)
}

func (uc *MediaUseCase) Update(media *entity.Media) error {
	return uc.repo.Update(media)
}

func (uc *MediaUseCase) Delete(id int64) error {
	return uc.repo.Delete(id)
}

func (uc *MediaUseCase) ListForDropdown(groupID int) ([]*entity.Media, error) {
	return uc.repo.ListForDropdown(groupID)
}
