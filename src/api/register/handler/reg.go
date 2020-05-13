package service

import (
	"context"
	"fmt"
	pb "grpc_blog/src/api/register/proto"
	"grpc_blog/src/api/register/service"
)

type Register struct {
}

func (h *Register) Call(ctx context.Context, in *pb.In) (*pb.Out, error) {
	s := &service.RegService{}
	s.Name = in.Name
	s.Pwd = in.Pwd
	s.Pwd2 = in.Pwds
	s.Email = in.Email
	s.Register()
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
