package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/errors"
	"back2/internal/domain/repository"

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
	return uc.repo.Create(staff)
}

func (uc *StaffUseCase) GetByID(id int64) (*entity.Staff, error) {
	return uc.repo.GetByID(id)
}

func (uc *StaffUseCase) GetByStaffID(staffID string) (*entity.Staff, error) {
	return uc.repo.GetByStaffID(staffID)
}

func (uc *StaffUseCase) List(groupID, page, pageSize int) ([]*entity.Staff, error) {
	offset := (page - 1) * pageSize
	return uc.repo.List(groupID, offset, pageSize)
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
	staff, err := uc.repo.GetByStaffID(staffID)
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
