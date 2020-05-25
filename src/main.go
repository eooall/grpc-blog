package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	articleHandle "grpc_blog/src/api/article/handle"
	articlePB "grpc_blog/src/api/article/proto"
	loginHandle "grpc_blog/src/api/login/handle"

	loginPB "grpc_blog/src/api/login/proto"
	regHandler "grpc_blog/src/api/register/handler"
	regPb "grpc_blog/src/api/register/proto"

	_ "grpc_blog/src/models"
	"net/http"
)

func main() {
	mux := runtime.NewServeMux()
	//登陆
	err := loginPB.RegisterLoginHandlerServer(context.TODO(), mux, &loginHandle.LoginService{})
	if err != nil {
		panic(err)
	}
	//注册
	err = regPb.RegisterRegisterHandlerServer(context.TODO(), mux, &regHandler.Register{})
	//添加文章
	err = articlePB.RegisterAddArticleHandlerServer(context.TODO(), mux, &articleHandle.AddArticle{})
	//获取文章
	err = articlePB.RegisterArticleListHandlerServer(context.TODO(), mux, &articleHandle.Article{})
	err = http.ListenAndServe(":8989", mux)
	if err != nil {
		panic(err)
	}
}
