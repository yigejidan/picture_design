package models

import (
	"errors"

	"gorm.io/gorm"

	"picture_design/common"
)

type Picture struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Name        string `gorm:"type:varchar(20);not null" json:"name"`
	User        string `gorm:"type:varchar(20);not null" json:"user"`
	Picture     string `gorm:"mediumblob;not null" json:"picture"`
	Description string `gorm:"varchar(255);not null" json:"description"`
	CreatedAt   int    `gorm:"int(11);not null" json:"created_at"`
	UpdatedAt   int    `gorm:"int(11);not null" json:"updated_at"`
	DeletedAt   int    `gorm:"int(11);not null" json:"deleted_at"`
}

type PictureData struct {
	List  []*Picture `json:"list"`
	Page  int64      `json:"page"`
	Size  int64      `json:"size"`
	Total int64      `json:"total"`
}

type PictureRepoImpl interface {
	Create(picture *Picture) error
	GetPicturesByUser(user string, page, size int) (*PictureData, error)
	GetPictureById(id int) (*PictureData, error)
	CheckNameIsDuplicate(user, name string) (bool, error)
}

type PictureRepo struct {
	db *gorm.DB
}

func NewPictureRepo() PictureRepoImpl {
	db := common.GetDB()
	pictureRepo := &PictureRepo{
		db: db,
	}
	return pictureRepo
}

func (u *PictureRepo) Create(picture *Picture) error {
	res := u.db.Create(picture)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *PictureRepo) GetPicturesByUser(user string, page, size int) (*PictureData, error) {
	var (
		pictures []*Picture
		total    int64
		offset   = (page - 1) * size
	)
	query := u.db.Table("pictures").Where("user = ? and name != ''", user)
	if res := query.Count(&total); res.Error != nil {
		return nil, query.Error
	}
	if res := query.Offset(offset).Limit(size).Find(&pictures); res.Error != nil {
		return nil, res.Error
	}
	pictureData := &PictureData{
		List:  pictures,
		Page:  int64(page),
		Size:  int64(size),
		Total: total,
	}
	return pictureData, nil
}

func (u *PictureRepo) GetPictureById(id int) (*PictureData, error) {
	var pictures []*Picture
	query := u.db.Table("pictures").Where("id = ?", id)
	if query.Error != nil {
		return nil, query.Error
	}
	query = query.Find(&pictures)
	if query.Error != nil {
		return nil, query.Error
	}
	pictureData := &PictureData{
		List:  pictures,
		Page:  1,
		Size:  10,
		Total: 1,
	}
	return pictureData, nil
}

type PictureName struct {
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func (u *PictureRepo) CheckNameIsDuplicate(user, name string) (bool, error) {
	var pictureName *PictureName
	query := u.db.Table("pictures")
	query.QueryFields = true
	tx := query.Where("user = ? and name = ?", user, name).First(&pictureName)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
