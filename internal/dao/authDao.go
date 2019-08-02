package dao

import "github.com/brucewang11/frame/internal/model"

type AuthDao struct {
	BaseDao
}


func (s *AuthDao)AddAuth(auth *model.Auth) error{
	return s.Create(auth)
}
