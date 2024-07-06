package service

import (
	"back/models"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-xorm/xorm"
)

var DBEngine *xorm.Engine

type CastLoginService struct {
	DB *xorm.Engine
}

func NewCastLoginService(db *xorm.Engine) *CastLoginService {
	return &CastLoginService{DB: db}
}

func (s *CastLoginService) generateCastUniqueRandomID() (string, error) {
	rand.Seed(time.Now().UnixNano()) // 乱数生成器を初期化

	for i := 0; i < 100; i++ { // 最大100回試行
		id := fmt.Sprintf("C%06d", rand.Intn(1000000)) // 'C'で始まる7桁のIDを生成
		exists, err := s.DB.Exist(&models.CastLogin{CastID: id})
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

func (s *CastLoginService) Create(castLogin *models.CastLogin) error {
	session := s.DB.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	// ランダムIDを生成
	id, err := s.generateCastUniqueRandomID()
	if err != nil {
		session.Rollback()
		return err
	}
	castLogin.CastID = id

	if castLogin.Status == "" {
		castLogin.Status = "active"
	}

	// デバッグ用のログを追加
	fmt.Printf("挿入前のCastLogin: %+v\n", castLogin)

	if _, err := session.Insert(castLogin); err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}

func (s *CastLoginService) GetByID(id string) (*models.CastLogin, error) {
	castLogin := &models.CastLogin{CastID: id}
	has, err := s.DB.Get(castLogin)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return castLogin, nil
}

func (s *CastLoginService) Update(castLogin *models.CastLogin) error {
	_, err := s.DB.ID(castLogin.ID).Update(castLogin)
	return err
}

func (s *CastLoginService) Delete(id int64) error {
	_, err := s.DB.ID(id).Delete(&models.CastLogin{})
	return err
}

func (s *CastLoginService) List(page, pageSize int) ([]*models.CastLogin, error) {
	castLogins := make([]*models.CastLogin, 0)
	err := s.DB.Limit(pageSize, (page-1)*pageSize).Find(&castLogins)
	return castLogins, err
}
