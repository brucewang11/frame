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


	}
	return &CodeModel{Code:0}
}

func UpdateAuth(auth *vo.AuthVo) *CodeModel{
	condition := model.Auth{
		Name:auth.Name,
	}

	userDao := &dao.AuthDao{BaseDao:dao.NewBaseDao(nil)}
	err := userDao.Update(condition,&model.Auth{AuthType:"wfwfwf1"})
	if err!=nil {
		fmt.Println(err)

	}
	return &CodeModel{Code:0}
}

func DeleteAuth(auth *vo.AuthVo) *CodeModel{
	condition := model.Auth{
		Name:auth.Name,
	}
	userDao := &dao.AuthDao{BaseDao:dao.NewBaseDao(nil)}
	err := userDao.Delete(&condition)
	if err!=nil {
		fmt.Println(err)
	}
	return &CodeModel{Code:0}
}

func ListAuth(auth *vo.AuthVo) *CodeModel{
	panic("1")
	list := []model.Auth{}
	condition := model.Auth{
		Name:auth.Name,
	}
	userDao := &dao.AuthDao{BaseDao:dao.NewBaseDao(nil)}
	err := userDao.List(condition,&list)
	if err!=nil {
		fmt.Println(err)
	}
	maps := make(map[string]interface{})
	maps["name"] = "wang"
	maps["num"] = "5"
	return &CodeModel{Code:1002,Data:list,TemplateData:maps}
}
