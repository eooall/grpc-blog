package tool

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

func InputCheck(t, v string) bool {
	switch t {
	//验证账户是否合法
	case "account":
		if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", v); !ok {
			return false
		}
		return true
	case "pwd":
		if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{6,16}$", v); !ok {
			return false
		}
		return true
	case "email":
		if ok, _ := regexp.MatchString("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$", v); !ok {
			return false
		}
		return true
	default:
		return false
	}

}

//md5生成
func NewMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
