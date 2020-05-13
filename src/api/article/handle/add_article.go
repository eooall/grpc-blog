package handle

import (
	"context"
	pb "grpc_blog/src/api/article/proto"
)

type AddArticle struct{}

func (h *AddArticle) Call(ctx context.Context, in *pb.AddArticleIn) (*pb.AddArticleOut, error) {

	return nil, nil
}
