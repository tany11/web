package repository

import "back2/internal/domain/entity"

type StaffRepository interface {
	Create(staff *entity.Staff) error
	GetByID(id int64) (*entity.Staff, error)
	GetByStaffIID(staffID string) (*entity.Staff, error)
	List(groupID int) ([]*entity.Staff, error)
	Update(staff *entity.Staff) error
	Delete(id int64) error
	GetByEmail(email string) (*entity.Staff, error)
	ListForDropdown(groupID int) ([]*entity.Staff, error) // 新しいメソッドを追加
}
