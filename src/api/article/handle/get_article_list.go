package handle

import (
	"context"
	pb "grpc_blog/src/api/article/proto"
)

type Article struct{}

func (h *Article) Call(ctx context.Context, in *pb.ArticleIn) (*pb.ArticleListOut, error) {

	return nil, nil
}
