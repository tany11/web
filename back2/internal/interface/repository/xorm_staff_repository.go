package repository

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"

	"github.com/go-xorm/xorm"
)

type XormStaffRepository struct {
	engine *xorm.Engine
}

func NewXormStaffRepository(engine *xorm.Engine) repository.StaffRepository {
	return &XormStaffRepository{engine: engine}
}

func (r *XormStaffRepository) Create(staff *entity.Staff) error {
	_, err := r.engine.Insert(staff)
	return err
}

func (r *XormStaffRepository) GetByID(id int64) (*entity.Staff, error) {
	staff := new(entity.Staff)
	has, err := r.engine.ID(id).Get(staff)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return staff, nil
}

func (r *XormStaffRepository) GetByStaffIID(staffID string) (*entity.Staff, error) {
	staff := new(entity.Staff)
	has, err := r.engine.Where("staff_i_d = ?", staffID).Get(staff)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return staff, nil
}

func (r *XormStaffRepository) List(groupID int) ([]*entity.Staff, error) {
	staffs := make([]*entity.Staff, 0)
	err := r.engine.Where("group_i_d = ? AND is_deleted = ?", groupID, "0").Find(&staffs)
	if err != nil {
		return nil, err
	}
	return staffs, nil
}

func (r *XormStaffRepository) Update(staff *entity.Staff) error {
	_, err := r.engine.ID(staff.ID).Update(staff)
	return err
}

func (r *XormStaffRepository) Delete(id int64) error {
	_, err := r.engine.ID(id).Delete(new(entity.Staff))
	return err
}

func (r *XormStaffRepository) GetByEmail(email string) (*entity.Staff, error) {
	staff := new(entity.Staff)
	has, err := r.engine.Where("email = ?", email).Get(staff)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return staff, nil
}

func (r *XormStaffRepository) ListForDropdown(groupID int) ([]*entity.Staff, error) {
	staffs := make([]*entity.Staff, 0)
	err := r.engine.Where("group_i_d = ?", groupID).Cols("i_d", "staff_i_d", "staff_last_name", "office_flg", "driver_flg").Find(&staffs)
	return staffs, err
}
