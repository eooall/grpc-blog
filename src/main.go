package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	LoginHandle "grpc_blog/src/api/login/handle"
	loginPB "grpc_blog/src/api/login/proto"
	regHandler "grpc_blog/src/api/register/handler"
	regPb "grpc_blog/src/api/register/proto"
	_ "grpc_blog/src/models"
	"net/http"
)

func main() {
	mux := runtime.NewServeMux()
	err := loginPB.RegisterLoginHandlerServer(context.TODO(), mux, &LoginHandle.LoginService{})
	if err != nil {
		panic(err)
	}
	err = regPb.RegisterRegisterHandlerServer(context.TODO(), mux, &regHandler.Register{})
	err = http.ListenAndServe(":8989", mux)
	if err != nil {
		panic(err)
	}
}
