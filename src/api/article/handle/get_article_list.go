package handle

import (
	"context"
	"fmt"
	pb "grpc_blog/src/api/article/proto"
	"grpc_blog/src/api/article/service"
)

type Article struct{}

func (h *Article) Call(ctx context.Context, in *pb.ArticleIn) (*pb.ArticleListOut, error) {
	s := &service.GetArticleList{}
	s.Search = in.Search
	s.ArticleType = in.ArticleType
	s.PageCurrent = int(in.PageCurrent)
	s.GetArticleList()
	if s.Err != nil {
		return &pb.ArticleListOut{
			Data:    &pb.ArticleListData{},
			Code:    int64(s.Code),
			Message: fmt.Sprintf("%s", s.Err),
		}, nil
	}

	return &pb.ArticleListOut{
		Data: &pb.ArticleListData{
			Article:   s.Data.Data,
			PageCount: s.Data.Total,
		},
		Code:    int64(s.Code),
		Message: s.Message,
	}, nil
}
