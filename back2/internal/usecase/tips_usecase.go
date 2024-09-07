package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
	"fmt"
	"time"
)

type TipsUseCase struct {
	tipsRepo repository.TipsRepository
}

func NewTipsUseCase(tipsRepo repository.TipsRepository) *TipsUseCase {
	return &TipsUseCase{tipsRepo: tipsRepo}
}

func (uc *TipsUseCase) Create(tips *entity.Tips) error {
	tips.CreatedAt = time.Now()
	tips.UpdatedAt = time.Now()

	err := uc.tipsRepo.Create(tips)
	if err != nil {
		return fmt.Errorf("ヒントの作成に失敗しました: %w", err)
	}

	return nil
}

func (uc *TipsUseCase) Update(tips *entity.Tips) error {
	tips.UpdatedAt = time.Now()
	return uc.tipsRepo.Update(tips)
}

func (uc *TipsUseCase) List(groupID, page, pageSize int) ([]*entity.Tips, error) {

	offset := (page - 1) * pageSize
	return uc.tipsRepo.List(groupID, offset, pageSize)
}

func (uc *TipsUseCase) ListReserved(groupID, page, pageSize int) ([]*entity.Tips, error) {
	offset := (page - 1) * pageSize
	return uc.tipsRepo.ListReserved(groupID, offset, pageSize)
}

func (uc *TipsUseCase) GetByID(id int64) (*entity.Tips, error) {
	return uc.tipsRepo.GetByID(id)
}

func (uc *TipsUseCase) Delete(id int64) error {
	return uc.tipsRepo.Delete(id)
}

func (uc *TipsUseCase) UpdateCompletionFlg(id int64) error {
	return uc.tipsRepo.UpdateCompletionFlg(id)
}

func (uc *TipsUseCase) UpdateIsDeleted(id int64) error {
	return uc.tipsRepo.UpdateIsDeleted(id)
}

func (uc *TipsUseCase) ListSchedule(groupID int, startDate, endDate string) ([]*entity.Tips, error) {
	return uc.tipsRepo.ListSchedule(groupID, startDate, endDate)
}
