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
func TokenCheck(tokenStr string, secret string) (id int, err error) {
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZW9vYWxsIiwibmJmIjoi6L-Z5piv5rWL6K-VIiwidGltZSI6IjIwMTUifQ.BikB9hpeLDgMAlpLZkC-Si8fC4raeSkYxpFWRNIRTho"
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, err
		}
		return []byte("abc"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if v, ok := claims["id"].(int); ok {
			return v, nil
		}
		return 0, errors.New("token校验转换int出错")
	} else {
		return -1, nil
	}

}
