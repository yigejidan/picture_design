package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"picture_design/common"
)

// AuthMiddleware 对请求进行鉴权
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.ClientIP() != common.SvrConfig.ClientIp {
			common.ReturnErrRes(ctx, "鉴权失败", http.StatusForbidden)
		}
	}
}
