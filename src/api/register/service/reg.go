package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"grpc_blog/src/api/register/models"
	"grpc_blog/src/tool"
	"time"
)

type RegService struct {
	Name    string
	Email   string
	Pwd     string
	Pwd2    string
	Token   string
	Code    int32
	id      int64
	Message string
	Err     error
}

func (s *RegService) Register() *RegService {
	s.check().add().awardToken()
	return s
}

//检查信息是否合法
func (s *RegService) check() *RegService {
	if !tool.InputCheck("account", s.Name) && !tool.InputCheck("email", s.Email) && !tool.InputCheck("pwd", s.Pwd) {
		s.Err = errors.New("账号或者密码格式不正确")
		s.Code = 403
		return s
	}
	if s.Pwd != s.Pwd2 {
		s.Err = errors.New("账号或者密码格式不正确")
		s.Code = 403
	}
	return s
}

//插入数据
func (s *RegService) add() *RegService {
	s.check()
	if s.Err != nil {
		return s
	}
	u := models.Users{}
	u.Email = s.Email
	u.Name = s.Name
	u.Password = tool.NewMd5(s.Pwd)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.Avatar = "http://a2.att.hudong.com/36/48/19300001357258133412489354717.jpg"
	id, err := models.AddUser(&u)
	if err != nil {
		s.Err = errors.New("服务器内部发生错误")
		s.Code = 500
		return s
	}
	s.id = id
	return s
}

//生成token
func (s *RegService) awardToken() *RegService {
	if s.Err != nil {
		return s
	}
	token, err := tool.TokenAward(jwt.MapClaims{
		"id": s.id,
	})
	if err != nil {
		s.Err = errors.New("token发布失败")
		s.Code = 500
		return s
	}
	s.Token = token
	s.Message = "登录成功"
	return s
}
