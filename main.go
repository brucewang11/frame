// Copyright 2019. bruce.wang authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		 http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


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