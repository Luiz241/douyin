package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func GetToken(secretKey string, iat, seconds int64, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenString, key string) (int64, error) {

	type MyClaims struct {
		jwt.RegisteredClaims
		UserId int64 `json:"userId"`
	}

	mytoken, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("错误的签名方法：%v", t.Header["alg"])
		}
		return []byte(key), nil
	})

	//mytoken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
	//	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("错误的签名方法：%v", t.Header["alg"])
	//	}
	//	return []byte(key), nil
	//})
	if err != nil {
		return 0, err
	}
	if claims, ok := mytoken.Claims.(*MyClaims); ok && mytoken.Valid {
		return claims.UserId, nil
	} else {
		return 0, fmt.Errorf("token失效")
	}

}
