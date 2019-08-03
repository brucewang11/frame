package main

import (
	log "github.com/alecthomas/log4go"
	"github.com/brucewang11/frame/common"
	"github.com/brucewang11/frame/internal/dao"
	"github.com/brucewang11/frame/internal/router"
	"github.com/gin-gonic/gin"
)

func main(){
	//init common
	common.InitCommon()
	//init db
	dao.InitDB()
	//init log
	log.LoadConfiguration("./static/log.xml")

	//init gin
	gin.SetMode(common.HttpInfo.GinMode)
	router := router.InitRouter()
	router.Run(common.HttpInfo.Port)

}