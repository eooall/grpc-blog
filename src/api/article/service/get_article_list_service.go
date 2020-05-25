package service

import (
	"grpc_blog/src/api/article/models"
	pb "grpc_blog/src/api/article/proto"
)

type GetArticleList struct {
	ArticleType string
	Search      string
	PageCurrent int
	PageCount   int
	Message     string
	Code        int
	Err         error
	Data        ArticleData
}
type ArticleData struct {
	Data  []*pb.ArticleData
	Total int64
}

func (s *GetArticleList) GetArticleList() *GetArticleList {
	if s.PageCount <= 5 {
		s.PageCount = 5
	}
	if s.PageCurrent >= 1 {
		s.PageCurrent = (s.PageCurrent - 1) * s.PageCount
	}
	d, err := models.GetArticle(s.PageCurrent, s.PageCount)
	if err != nil {
		s.Err = err
		s.Code = 500
		return s
	}
	s.Data.Data = d
	total, err := models.GetTotal()
	if err != nil {
		s.Err = err
		s.Code = 500
		return s
	}
	s.Data.Total = total
	s.Code = 200
	s.Message = "查询成功"
	return s
}
