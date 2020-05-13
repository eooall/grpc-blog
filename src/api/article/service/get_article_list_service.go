package service

type GetArticleList struct {
	ArticleType string //类型
	Search      string //搜索关键字
	PageCurrent int64  //当前页面
}

func (s *GetArticleList) GetArticleList() *GetArticleList {

	return s
}
