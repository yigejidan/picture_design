package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ReturnErrRes(ctx *gin.Context, msg string, statusCode int) {
	ctx.JSON(statusCode, &Response{
		Code: 0,
		Data: []int{},
		Msg:  msg,
	})
}

func ReturnSuccessRes(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Code: 1,
		Data: data,
		Msg:  msg,
	})
}
