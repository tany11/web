package service

import (
	"back/models"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-xorm/xorm"
)

type StaffLoginService struct {
	DB *xorm.Engine
}

func NewStaffLoginService(db *xorm.Engine) *StaffLoginService {
	return &StaffLoginService{DB: db}
}

func (s *StaffLoginService) generateStaffUniqueRandomID() (string, error) {
	rand.Seed(time.Now().UnixNano()) // 乱数生成器を初期化

	for i := 0; i < 100; i++ { // 最大100回試行
		id := fmt.Sprintf("S%06d", rand.Intn(1000000)) // 'S'で始まる7桁のIDを生成
		exists, err := s.DB.Exist(&models.StaffLogin{StaffID: id})
		if err != nil {
			log.Printf("IDの存在確認中にエラーが発生: %v\n", err)
			return "", err
		}
		if !exists {
			log.Printf("生成されたユニークID: %s\n", id)
			return id, nil
		}
	}
	log.Println("ユニークなIDの生成に失敗しました")
	return "", fmt.Errorf("ユニークなIDの生成に失敗しました")
}

func (s *StaffLoginService) Create(staffLogin *models.StaffLogin) error {
	session := s.DB.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	// ランダムIDを生成
	id, err := s.generateStaffUniqueRandomID()
	if err != nil {
		session.Rollback()
		return err
	}
	staffLogin.StaffID = id

	if staffLogin.Status == "" {
		staffLogin.Status = "active"
	}

	// デバッグ用のログを追加
	fmt.Printf("挿入前のStaffLogin: %+v\n", staffLogin)

	if _, err := session.Insert(staffLogin); err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}

func (s *StaffLoginService) GetByID(id string) (*models.StaffLogin, error) {
	staffLogin := &models.StaffLogin{StaffID: id}
	has, err := s.DB.Get(staffLogin)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return staffLogin, nil
}

func (s *StaffLoginService) Update(staffLogin *models.StaffLogin) error {
	_, err := s.DB.ID(staffLogin.ID).Update(staffLogin)
	return err
}

func (s *StaffLoginService) Delete(id int64) error {
	_, err := s.DB.ID(id).Delete(&models.StaffLogin{})
	return err
}

func (s *StaffLoginService) List(page, pageSize int) ([]*models.StaffLogin, error) {
	staffLogins := make([]*models.StaffLogin, 0)
	err := s.DB.Limit(pageSize, (page-1)*pageSize).Find(&staffLogins)
	return staffLogins, err
}
