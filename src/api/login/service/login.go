package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"grpc_blog/src/api/login/models"
	"grpc_blog/src/tool"
)

type LoginService struct {
	Email   string
	Pwd     string
	Code    int32
	Err     error
	Message string
	Token   string
	id      int64
}

func (s *LoginService) Login() *LoginService {
	return s.check().checkMd5().awardToken()
}

func (s *LoginService) check() *LoginService {
	if !(tool.InputCheck("email", s.Email) && tool.InputCheck("pwd", s.Pwd)) {
		s.Code = 403
		s.Err = errors.New("账号或密码格式不正确")
		return s
	}
	return s
}

func (s *LoginService) checkMd5() *LoginService {
	if s.Err != nil {
		return s
	}
	u, err := models.GetUsersByEmail(s.Email)
	if err != nil {
		s.Code = 403
		s.Err = errors.New("当前账号不存在")
		return s
	}
	if !(tool.NewMd5(s.Pwd) == u.Password) {
		s.Code = 403
		s.Err = errors.New("账号或密码不正确")
		return s
	}
	s.id = u.ID
	return s
}

func (s *LoginService) awardToken() *LoginService {
	if s.Err != nil {
		return s
	}
	token, err := tool.TokenAward(jwt.MapClaims{
		"id": s.id,
	})
	if err != nil {
		s.Code = 500
		s.Err = errors.New("token颁发失败")
		return s
	}
	s.Token = token
	s.Message = "登录成功"
	return s
}
