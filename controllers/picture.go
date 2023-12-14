package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"picture_design/common"
	"picture_design/models"
)

type PictureImpl interface {
	SavePictures(ctx *gin.Context)
	GetPictures(ctx *gin.Context)
}

type PictureController struct {
}

type savePicturesForm struct {
	Name        string `form:"name"`
	Picture     string `form:"picture" binding:"required"`
	Description string `form:"description"`
}

func NewPicture() PictureImpl {
	picture := &PictureController{}
	return picture
}

func (p *PictureController) SavePictures(ctx *gin.Context) {
	authUser := models.CheckUserExist(ctx)
	if authUser == nil {
		return
	}
	json := savePicturesForm{}
	err := ctx.BindJSON(&json)
	if err != nil {
		common.Log("%v", "SavePictures params,err: "+err.Error())
		common.ReturnErrRes(ctx, "参数错误", http.StatusBadRequest)
		return
	}
	if json.Name != "" {
		isDuplicate, err := models.NewPictureRepo().CheckNameIsDuplicate(authUser.Name, json.Name)
		if err != nil {
			common.Log("%v", "SavePictures db GetPictureByName,err: "+err.Error())
			common.ReturnErrRes(ctx, "请求失败", http.StatusInternalServerError)
			return
		}
		if isDuplicate {
			common.Log("%v", "SavePictures name duplicate")
			common.ReturnErrRes(ctx, "效果图名称重复", http.StatusBadRequest)
			return
		}
	}
	newPicture := models.Picture{
		User:        authUser.Name,
		Name:        json.Name,
		Picture:     json.Picture,
		Description: json.Description,
	}
	err = models.NewPictureRepo().Create(&newPicture)
	if err != nil {
		common.Log("%v", "SavePictures db Create,err: "+err.Error())
		common.ReturnErrRes(ctx, "请求失败", http.StatusInternalServerError)
		return
	}
	common.ReturnSuccessRes(ctx, "效果图上传成功", map[string]uint{"id": newPicture.ID})
}

func (p *PictureController) GetPictures(ctx *gin.Context) {
	authUser := models.CheckUserExist(ctx)
	if authUser == nil {
		return
	}
	id, err := strconv.Atoi(ctx.DefaultQuery("id", "0"))
	if err != nil {
		common.Log("%v", "GetPictures id is not int,err: "+err.Error())
		common.ReturnErrRes(ctx, "参数错误", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		common.Log("%v", "GetPictures page is not int,err: "+err.Error())
		common.ReturnErrRes(ctx, "参数错误", http.StatusBadRequest)
		return
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err != nil {
		common.Log("%v", "GetPictures size is not int,err: "+err.Error())
		common.ReturnErrRes(ctx, "参数错误", http.StatusBadRequest)
		return
	}
	var pictures *models.PictureData
	if id != 0 {
		pictures, err = models.NewPictureRepo().GetPictureById(id)
		if err != nil {
			common.Log("%v", "GetPictures db GetPictureByName,err: "+err.Error())
			common.ReturnErrRes(ctx, "请求失败", http.StatusInternalServerError)
			return
		}
	} else {
		pictures, err = models.NewPictureRepo().GetPicturesByUser(authUser.Name, page, size)
		if err != nil {
			common.Log("%v", "GetPictures db GetPicturesByUser,err: "+err.Error())
			common.ReturnErrRes(ctx, "请求失败", http.StatusInternalServerError)
			return
		}
	}
	common.ReturnSuccessRes(ctx, "效果图下载成功", pictures)
}
