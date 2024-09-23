package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type CastUseCase struct {
	repo repository.CastRepository
}

func NewCastUseCase(repo repository.CastRepository) *CastUseCase {
	return &CastUseCase{repo: repo}
}

func (uc *CastUseCase) Create(cast *entity.Cast) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cast.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	cast.PasswordHash = string(hashedPassword)
	cast.GroupID = 1

	// Generate unique cast ID
	castID, err := uc.generateCastUniqueRandomID()
	if err != nil {
		return err
	}
	cast.CastID = castID

	return uc.repo.Create(cast)
}

func (uc *CastUseCase) GetByID(id int64) (*entity.Cast, error) {
	return uc.repo.GetByID(id)
}

func (uc *CastUseCase) List(groupID int) ([]*entity.Cast, error) {
	return uc.repo.List(groupID)
}

func (uc *CastUseCase) ListForDropdown(groupID int) ([]*entity.Cast, error) {
	casts, err := uc.repo.ListForDropdown(groupID)
	if err != nil {
		log.Printf("Error in CastUseCase.ListForDropdown: %v", err) // Added error logging
		return nil, err
	}
	return casts, nil
}

func (uc *CastUseCase) Update(cast *entity.Cast) error {
	return uc.repo.Update(cast)
}

func (uc *CastUseCase) Delete(id int64) error {
	return uc.repo.Delete(id)
}

func (uc *CastUseCase) generateCastUniqueRandomID() (string, error) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		id := fmt.Sprintf("C%06d", rand.Intn(1000000))
		cast, err := uc.repo.GetByCastIID(id)
		if err != nil {
			log.Printf("IDの存在確認中にエラーが発生: %v\n", err)
			return "", fmt.Errorf("IDの存在確認中にエラーが発生: %w", err)
		}
		if cast == nil {
			log.Printf("生成されたユニークID: %s\n", id)
			return id, nil
		}
	}
	log.Println("ユニークなIDの生成に失敗しました")
	return "", fmt.Errorf("ユニークなIDの生成に失敗しました")
}

func (uc *CastUseCase) UpdateWorkingFlg(id int64) error {
	existingCast, err := uc.repo.GetByID(id)
	if err != nil {
		log.Printf("IDでキャストを取得中にエラーが発生: %v\n", err)
		return fmt.Errorf("IDでキャストを取得中にエラーが発生: %w", err)
	}
	if existingCast == nil {
		log.Println("指定されたIDのキャストが見つかりません")
		return fmt.Errorf("指定されたIDのキャストが見つかりません")
	}

	// workingFlgをスイッチ
	if existingCast.WorkingFlg == "1" {
		existingCast.WorkingFlg = "0"
	} else {
		existingCast.WorkingFlg = "1"
	}

	return uc.repo.Update(existingCast)
}

func (uc *CastUseCase) ResetAllWorkingFlags() error {
	return uc.repo.ResetAllWorkingFlags()
}
