package model

type Auth struct {
	ID       int    `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"column:name"`
	AuthType string `gorm:"column:authType"`
}

func (auth Auth) TableName() string {
	return "auths"
}
