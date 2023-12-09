package models

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"picture_design/common"
)

type User struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
	Password  string `gorm:"varchar(255);not null" json:"password"`
	Type      int    `gorm:"int(10);not null" json:"type"`
	CreatedAt int    `gorm:"int(11);not null" json:"created_at"`
	UpdatedAt int    `gorm:"int(11);not null" json:"updated_at"`
	DeletedAt int    `gorm:"int(11);not null" json:"deleted_at"`
}

type UserRepoImpl interface {
	Create(user *User) error
	GetUser(name string) (*User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo() UserRepoImpl {
	db := common.GetDB()
	userRepo := &UserRepo{
		db: db,
	}
	return userRepo
}

func (u *UserRepo) Create(user *User) error {
	res := u.db.Create(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserRepo) GetUser(name string) (*User, error) {
	var user User
	res := u.db.Where("name = ?", name).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func CheckUserExist(ctx *gin.Context) *User {
	authUser := ctx.GetHeader("user")
	userRepo := NewUserRepo()
	dbUser, err := userRepo.GetUser(authUser)
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		common.Log("%v", "CheckUserExist db GetUser,err: "+err.Error())
		common.ReturnErrRes(ctx, "请求失败", http.StatusInternalServerError)
		return nil
	}
	if errors.Is(gorm.ErrRecordNotFound, err) {
		common.Log("%v", "CheckUserExist GetUser ErrRecordNotFound,err: "+err.Error())
		common.ReturnErrRes(ctx, "账号不存在", http.StatusForbidden)
		return nil
	}
	return dbUser
}
