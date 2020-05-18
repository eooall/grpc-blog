package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	addHandle "grpc_blog/src/api/article/handle"
	articlePB "grpc_blog/src/api/article/proto"
	LoginHandle "grpc_blog/src/api/login/handle"

	loginPB "grpc_blog/src/api/login/proto"
	regHandler "grpc_blog/src/api/register/handler"
	regPb "grpc_blog/src/api/register/proto"
	_ "grpc_blog/src/models"
	"net/http"
)

func main() {
	mux := runtime.NewServeMux()
	//登陆
	err := loginPB.RegisterLoginHandlerServer(context.TODO(), mux, &LoginHandle.LoginService{})
	if err != nil {
		panic(err)
	}
	//注册
	err = regPb.RegisterRegisterHandlerServer(context.TODO(), mux, &regHandler.Register{})
	//添加文章
	err = articlePB.RegisterAddArticleHandlerServer(context.TODO(), mux, &addHandle.AddArticle{})
	err = http.ListenAndServe(":8989", mux)
	if err != nil {
		panic(err)
	}
}
