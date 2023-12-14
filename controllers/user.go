package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"picture_design/common"
	"picture_design/models"
)

type UserImpl interface {
	CreateUser(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type loginForm struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type userForm struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
	Type     int    `form:"type"`
}

type UserController struct {
}

func NewUser() UserImpl {
	user := &UserController{}
	return user
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	authUser := models.CheckUserExist(ctx)
	if authUser == nil {
		return
	}
	if authUser.Type != 1 {
		common.ReturnErrRes(ctx, "非管理员不能创建账号", http.StatusForbidden)
		return
	}
	json := userForm{}
	err := ctx.BindJSON(&json)
	if err != nil {
		common.Log("%v", "CreateUser params,err: "+err.Error())
		common.ReturnErrRes(ctx, "参数错误", http.StatusBadRequest)
		return
	}
	if err != nil {
		common.Log("%v", "CreateUser userType is not int,err: "+err.Error())
		common.ReturnErrRes(ctx, "参数错误", http.StatusBadRequest)
		return
	}
	newUser := models.User{
		Name:     json.Name,
		Password: json.Password,
		Type:     json.Type,
	}
	err = models.NewUserRepo().Create(&newUser)
	if err != nil {
		if err.Error() == "Error 1062 (23000): Duplicate entry 'test' for key 'name_index'" {
			common.Log("%v", "CreateUser name duplicate,err: "+err.Error())
			common.ReturnErrRes(ctx, "账号名称重复", http.StatusBadRequest)
			return
		}
		common.Log("%v", "CreateUser db Create,err: "+err.Error())
		common.ReturnErrRes(ctx, "账号创建失败", http.StatusInternalServerError)
		return
	}
	common.ReturnSuccessRes(ctx, "账号创建成功", []int{})
}

func (u *UserController) Login(ctx *gin.Context) {
	json := loginForm{}
	err := ctx.BindJSON(&json)
	if err != nil {
		common.Log("%v", "CreateUser params,err: "+err.Error())
		common.ReturnErrRes(ctx, "参数错误", http.StatusBadRequest)
		return
	}
	userRepo := models.NewUserRepo()
	dbUser, err := userRepo.GetUser(json.Name)
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		common.Log("%v", "CheckUserExist db GetUser,err: "+err.Error())
		common.ReturnErrRes(ctx, "请求失败", http.StatusInternalServerError)
		return
	}
	if errors.Is(gorm.ErrRecordNotFound, err) {
		common.Log("%v", "CheckUserExist GetUser ErrRecordNotFound,err: "+err.Error())
		common.ReturnErrRes(ctx, "账号不存在", http.StatusForbidden)
		return
	}
	if dbUser.Password != json.Password {
		common.Log("%v", "CreateUser password")
		common.ReturnErrRes(ctx, "密码错误", http.StatusForbidden)
		return
	}
	common.ReturnSuccessRes(ctx, "登陆成功", []int{})
}
