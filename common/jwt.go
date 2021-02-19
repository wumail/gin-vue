package common

import (
	"main/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_crect")

//Claims 参数
type Claims struct {
	//UserId d
	UserID uint
	jwt.StandardClaims
}

//ReleaseToken 加密
func ReleaseToken(user model.User) (string, error) {
	exprationTime := time.Now().Add(1 * 24 * time.Hour)
	//token内容
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exprationTime.Unix(),  //过期时间
			IssuedAt:  time.Now().Unix(),     //发放时间
			Issuer:    "1world1dream.wumail", //发放机构
			Subject:   "user token",          //主题
		},
	}
	//使用密钥生成token
	//加密内容和加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, err
}

//ParseToken 解密
func ParseToken(tokenStrng string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStrng, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err

}
