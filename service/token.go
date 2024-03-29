package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 自定义令牌
var mySigningKey = []byte("Key of ljp")

// MyClaim 对每个用户发放的权限
type MyClaim struct {
	Username       interface{}        // 用户名
	Id             int                // id
	StandardClaims jwt.StandardClaims // 使用jwt实现鉴权
}

// Valid 有效性
func (m MyClaim) Valid() error {
	return nil
}

// CreateToken 创建token
func CreateToken(userid int, username interface{}) (s string, err error) {
	// Create the Claims
	claims := MyClaim{
		Username: username,
		Id:       userid,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,    //生效时间，这里是一分钟前生效
			ExpiresAt: time.Now().Unix() + 60*60, //过期时间，这里是一小时过期
			Issuer:    "ljp",                     //签发人
		},
	}
	//SigningMethodHS256,HS256对称加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//通过自定义令牌加密
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("生成token失败")
	}
	return ss, err
}
