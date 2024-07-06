package repository

import "back2/internal/domain/entity"

type StaffRepository interface {
	Create(staff *entity.Staff) error
	GetByID(id int64) (*entity.Staff, error)
	GetByStaffID(staffID string) (*entity.Staff, error)
	List(groupID, offset, limit int) ([]*entity.Staff, error)
	Update(staff *entity.Staff) error
	Delete(id int64) error
}
