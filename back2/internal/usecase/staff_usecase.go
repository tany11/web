package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/errors"
	"back2/internal/domain/repository"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type StaffUseCase struct {
	repo repository.StaffRepository
}

func NewStaffUseCase(repo repository.StaffRepository) *StaffUseCase {
	return &StaffUseCase{repo: repo}
}

func (uc *StaffUseCase) Create(staff *entity.Staff) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(staff.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	staff.PasswordHash = string(hashedPassword)
	staff.GroupID = 1

	// Generate unique staff ID
	staffID, err := uc.generateStaffUniqueRandomID()
	if err != nil {
		return err
	}
	staff.StaffID = staffID

	return uc.repo.Create(staff)
}

func (uc *StaffUseCase) GetByID(id int64) (*entity.Staff, error) {
	return uc.repo.GetByID(id)
}

func (uc *StaffUseCase) GetByStaffIID(staffID string) (*entity.Staff, error) {
	return uc.repo.GetByStaffIID(staffID)
}

func (uc *StaffUseCase) List(groupID, page, pageSize int) ([]*entity.Staff, error) {
	offset := (page - 1) * pageSize
	return uc.repo.List(groupID, offset, pageSize)
}

func (uc *StaffUseCase) ListForDropdown(groupID int) ([]*entity.Staff, error) {
	staffs, err := uc.repo.ListForDropdown(groupID)
	if err != nil {
		log.Printf("Error in StaffUseCase.ListForDropdown: %v", err) // Added error logging
		return nil, err
	}
	return staffs, nil
}

func (uc *StaffUseCase) Update(staff *entity.Staff) error {
	if staff.PasswordHash != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(staff.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		staff.PasswordHash = string(hashedPassword)
	}
	return uc.repo.Update(staff)
}

func (uc *StaffUseCase) Delete(id int64) error {
	return uc.repo.Delete(id)
}

func (uc *StaffUseCase) Authenticate(staffID, password string) (*entity.Staff, error) {
	staff, err := uc.repo.GetByStaffIID(staffID)
	if err != nil {
		return nil, err
	}
	if staff == nil {
		return nil, errors.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(staff.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.ErrInvalidCredentials
	}

	return staff, nil
}

func (uc *StaffUseCase) GetByEmail(email string) (*entity.Staff, error) {
	return uc.repo.GetByEmail(email)
}

func (uc *StaffUseCase) generateStaffUniqueRandomID() (string, error) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		id := fmt.Sprintf("S%06d", rand.Intn(1000000))
		staff, err := uc.repo.GetByStaffIID(id)
		if err != nil {
			log.Printf("ID存在確認中にエラーが発生: %v\n", err)
			return "", fmt.Errorf("IDの存在確認中にエラーが発生: %w", err)
		}
		if staff == nil {
			log.Printf("生成されたユニークID: %s\n", id)
			return id, nil
		}
	}
	log.Println("ユニークなIDの生成に失敗しました")
	return "", fmt.Errorf("ユニークなIDの生成に失敗しました")
}
