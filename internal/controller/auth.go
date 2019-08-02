package controller

import (
	"github.com/brucewang11/frame/internal/service"
	"github.com/brucewang11/frame/internal/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddAuth(ctx *gin.Context) {
	var auth vo.AuthVo
	if err := ctx.ShouldBind(&auth); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"Code": 1001, "Msg":"参数错误"})
	}

	res := service.AddAuth(&auth)
	ResData(res,ctx)
}
