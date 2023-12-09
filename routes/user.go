package routes

import (
	"github.com/gin-gonic/gin"
	"picture_design/controllers"
)

func UserRoute(r *gin.Engine) *gin.Engine {
	user := controller.NewUser()
	r.POST("/users", user.CreateUser)
	r.POST("/login", user.Login)
	return r
}
