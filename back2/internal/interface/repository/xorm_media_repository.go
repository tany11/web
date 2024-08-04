package repository

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
	"log"

	"github.com/go-xorm/xorm"
)

type XormMediaRepository struct {
	engine *xorm.Engine
}

func NewXormMediaRepository(engine *xorm.Engine) repository.MediaRepository {
	return &XormMediaRepository{engine: engine}
}

func (r *XormMediaRepository) Create(media *entity.Media) error {
	_, err := r.engine.Insert(media)
	return err
}

func (r *XormMediaRepository) GetByID(id int64) (*entity.Media, error) {
	media := new(entity.Media)
	has, err := r.engine.ID(id).Get(media)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return media, nil
}

func (r *XormMediaRepository) GetByUsername(username string) (*entity.Media, error) {
	media := new(entity.Media)
	has, err := r.engine.Where("username = ?", username).Get(media)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return media, nil
}

func (r *XormMediaRepository) List(offset, limit int) ([]*entity.Media, error) {
	medias := make([]*entity.Media, 0)
	err := r.engine.Limit(limit, offset).Find(&medias)

	return medias, err
}

func (r *XormMediaRepository) Update(media *entity.Media) error {
	_, err := r.engine.ID(media.ID).Update(media)
	return err
}

func (r *XormMediaRepository) Delete(id int64) error {
	_, err := r.engine.ID(id).Delete(new(entity.Media))
	return err
}

func (r *XormMediaRepository) ListForDropdown(groupID int) ([]*entity.Media, error) {
	log.Println("ListForDropdown")
	medias := make([]*entity.Media, 0)
	err := r.engine.Cols("i_d", "media_name").Find(&medias)
	return medias, err
}
