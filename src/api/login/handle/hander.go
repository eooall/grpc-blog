package handle

import (
	"context"
	"fmt"
	pb "grpc_blog/src/api/login/proto"
	"grpc_blog/src/api/login/service"
)

type LoginService struct{}

func (h *LoginService) Call(ctx context.Context, in *pb.Input) (*pb.Out, error) {
	s := &service.LoginService{}
	s.Email = in.Email
	s.Pwd = in.Pwd
	s.Login()
	if s.Err != nil {
		return &pb.Out{
			Code:    s.Code,
			Data:    &pb.Data{},
			Message: fmt.Sprintf("%s", s.Err),
		}, nil
	}
	return &pb.Out{
		Code: 200,
		Data: &pb.Data{
			Token: s.Token,
		},
		Message: s.Message,
	}, nil
}
