package vo


type AuthVo struct{
	ID int  `form:"id"`
	Name string `form:"name"`
	AuthType string `form:"auth_type"`
}
