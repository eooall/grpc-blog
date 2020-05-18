package handle

import (
	"context"
	"fmt"
	pb "grpc_blog/src/api/article/proto"
	"grpc_blog/src/api/article/service"
)

type AddArticle struct{}

func (h *AddArticle) Call(ctx context.Context, in *pb.AddArticleIn) (*pb.AddArticleOut, error) {
	s := &service.AddArticle{}
	s.Title = in.Title
	s.Body = in.Body
	s.CategoryId = in.CategoryId
	s.Slug = in.Slug
	s.Excerpt = in.Excerpt
	s.Token = in.Token
	if s.AddArticle(); s.Err != nil {
		return &pb.AddArticleOut{
			Code:    s.Code,
			Message: fmt.Sprintf("%s", s.Err),
			Data:    &pb.OutData{},
		}, nil
	}
	return &pb.AddArticleOut{
		Code:    200,
		Message: s.Message,
		Data:    &pb.OutData{},
	}, nil
}
