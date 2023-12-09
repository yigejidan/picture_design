package routes

import (
	"github.com/gin-gonic/gin"
	controller "picture_design/controllers"
)

func PictureRoute(r *gin.Engine) *gin.Engine {
	picture := controller.NewPicture()
	r.POST("/pictures", picture.SavePictures)
	r.GET("/pictures", picture.GetPictures)
	return r
}
