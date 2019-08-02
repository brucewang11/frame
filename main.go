package main

import (
	"github.com/brucewang11/frame/common"
	"github.com/brucewang11/frame/internal/dao"
	"github.com/brucewang11/frame/internal/router"
	"github.com/brucewang11/frame/utils"
	log "github.com/alecthomas/log4go"
	"github.com/gin-gonic/gin"
	"path"
)

func main(){
	//init common
	common.InitCommon()
	//init db
	dao.InitDB()
	//init log
	logFile := path.Join(utils.GetCurrentDirectory(), "static/log.xml")
	log.LoadConfiguration(logFile)

	//init gin
	gin.SetMode(common.HttpInfo.GinMode)
	router := router.InitRouter()
	router.Run(common.HttpInfo.Port)

}