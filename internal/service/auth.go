package service

import (
	"fmt"
	"github.com/brucewang11/frame/internal/dao"
	"github.com/brucewang11/frame/internal/model"
	"github.com/brucewang11/frame/internal/vo"
)




func AddAuth(auth *vo.AuthVo) *CodeModel{
	param := model.Auth{
		Name:auth.Name,
		AuthType:auth.AuthType,
	}

	userDao := &dao.AuthDao{BaseDao:dao.NewBaseDao(nil)}
	err := userDao.Create(&param)
	if err!=nil {
		fmt.Println(err)

	}
	return &CodeModel{Code:0}
}
