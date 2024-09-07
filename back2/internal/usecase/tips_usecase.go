package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/errors"
	"back2/internal/domain/repository"
	"fmt"
	"log"
	"time"
)

type TipsUseCase struct {
	tipsRepo    repository.TipsRepository
}

func NewTipsUseCase(tipsRepo repository.TipsRepository) *TipsUseCase {
	return &TipsUseCase{
		tipsRepo: tipsRepo,
	}
}

func (uc *TipsUseCase) Create(tips *entity.Tips) error 

	tips.CreatedAt = time.Now()
	tips.UpdatedAt = time.Now()

	// 注文を作成
	err = uc.tipsRepo.Create(tips)
	if err != nil {
		log.Printf("ヒントの作成に失敗しました: %v", err)
		return fmt.Errorf("ヒントの作成に失敗しました: %w", err)
	}

	log.Printf("ヒントを作成しました: ID=%d", tips.ID)
	return nil
}
func (u *TipsUseCase) Update(tips *entity.Tips) error {
	return u.repository.Update(tips)
}

func (uc *TipsUseCase) List(groupID, page, pageSize int) ([]*entity.Tips, error) {
	offset := (page - 1) * pageSize
	tips, err := uc.tipsRepo.List(groupID, offset, pageSize)
	if err != nil {
		log.Printf("注文リストの取得中にエラーが発生しました: %v", err)
		return nil, fmt.Errorf("注文リストの取得に失敗しました: %w", err)
	}
	return tips, nil
}

func (uc *TipsUseCase) ListReserved(groupID, page, pageSize int) ([]*entity.Tips, error) {
	offset := (page - 1) * pageSize
	tips, err := uc.tipsRepo.ListReserved(groupID, offset, pageSize)
	if err != nil {
		log.Printf("予約リストの取得中にエラーが発生しました: %v", err)
		return nil, fmt.Errorf("予約リストの取得に失敗しました: %w", err)
	}
	return tips, nil
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

func (uc *TipsUseCase) ListSchedule(startDate, endDate string) ([]*entity.Tips, error) {
	return uc.tipsRepo.ListSchedule(startDate, endDate)
}

func (uc *TipsUseCase) Update(tips *entity.Tips) error {
	return uc.tipsRepo.Update(tips)
}
