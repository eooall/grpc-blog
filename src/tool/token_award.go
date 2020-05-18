package tool

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

//颁发token
func TokenAward(MapClaims jwt.MapClaims) (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MapClaims)
	tokenStr, err = token.SignedString([]byte("abc"))
	return
}

//校验token,返回用户ID
//-1表示token校验出错
//secret string
func TokenCheck(tokenStr string) (id float64, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, err
		}
		return []byte("abc"), nil
	})
	if err != nil {
		return -1, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if v, ok := claims["id"].(float64); ok {
			return v, nil
		}
		return 0, errors.New("token校验转换float64出错")
	} else {
		return -1, nil
	}

}
