package router

import (
	"errors"
	log "github.com/alecthomas/log4go"
	"github.com/brucewang11/frame/internal/controller"
	"github.com/brucewang11/frame/internal/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

const TOKEN_EXP = 24 // token有效周期

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		//非法token定义
		msgtoken := service.CodeModel{Code: 1001}
		if strings.ToUpper(c.Request.Method) == "GET" {
			token = c.Query("token")
		}
		if strings.ToUpper(c.Request.Method) == "POST" {
			token = c.PostForm("token")
		}
		if token == "" {
			log.Error("JWTAuth token is null")
			c.Abort()
			controller.ResData(&msgtoken,c)
			return
		}

		j := NewJWT()
		claims, err := j.ParseToken(token)
		if claims == nil {
			log.Error("JWTAuth token parse error")
			c.Abort()
			controller.ResData(&msgtoken,c)
			return
		}
		if err != nil {
			c.Abort()
			log.Error("token parse error", err.Error())
			controller.ResData(&msgtoken,c)
			return
		}
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "box.la"
)

type CustomClaims struct {
	ID       int
	AppID    string `json:"appid"`
	Account  string `json:"account"`
	UserType int    `json:"userType"`
	PubKey   string
	Level    int
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}
func GetSignKey() string {
	return SignKey
}
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil || token.Valid == false {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, TokenInvalid
}

//func (j *JWT) RefreshToken(tokenString string) (string, error) {
//
//	jwt.TimeFunc = func() time.Time {
//		return time.Unix(0, 0)
//	}
//	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return j.SigningKey, nil
//	})
//	if err != nil {
//		return "", err
//	}
//	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
//		jwt.TimeFunc = time.Now
//		claims.StandardClaims.ExpiresAt = time.Now().Add(TOKEN_EXP * time.Hour).Unix()
//		return j.CreateToken(*claims)
//	}
//	return "", TokenInvalid
//}


