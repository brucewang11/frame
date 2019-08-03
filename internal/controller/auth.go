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


func UpdateAuth(ctx *gin.Context) {
	var auth vo.AuthVo
	if err := ctx.ShouldBind(&auth); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	} else {
		res := service.UpdateAuth(&auth)
		ResData(res,ctx)
	}
}

func DelAuth(ctx *gin.Context) {
	var auth vo.AuthVo
	if err := ctx.ShouldBind(&auth); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	}
	res := service.DeleteAuth(&auth)
	ResData(res,ctx)
}

func ListAuth(ctx *gin.Context) {
	var auth vo.AuthVo
	if err := ctx.ShouldBind(&auth); err != nil {
		ResData(&service.CodeModel{Code:1001},ctx)
	}
	res := service.ListAuth(&auth)
	ResData(res,ctx)
}


