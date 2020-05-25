package service

import (
	"errors"
	"grpc_blog/src/api/article/models"
	"grpc_blog/src/tool"
	"time"
)

type AddArticle struct {
	Token      string
	Body       string
	CategoryId int64  //类别id
	Excerpt    string //摘要
	Slug       string //友好SEO url
	id         int64  //当前用户的id
	Title      string
	Code       int64
	Message    string
	Err        error
}

func (s *AddArticle) AddArticle() *AddArticle {
	id, err := tool.TokenCheck(s.Token)
	if err != nil {
		s.Code = 403
		s.Err = errors.New("token验证失败")
		return s
	}
	m := &models.AddArticleStruct{}
	m.UpdatedAt = time.Now()
	m.CreateTime = time.Now()
	m.Body = s.Body
	m.Slug = s.Slug
	m.Excerpt = s.Excerpt
	m.Title = s.Body
	m.UserId = int64(id)
	m.CategoryId = s.CategoryId
	err = models.AddArticle(m)
	if err != nil {
		s.Code = 500
		s.Err = errors.New("插入失败")
		return s
	} else {
		s.Code = 200
		s.Message = "添加成功"
	}
	return s
}
