package models

import (
	"grpc_blog/src/models"
	"time"
)

type AddArticleStruct struct {
	Body       string
	CategoryId int64 //类别id
	UserId     int64
	Excerpt    string //摘要
	Title      string
	Slug       string //友好SEO url
	CreateTime time.Time
	UpdatedAt  time.Time
}

func AddArticle(m *AddArticleStruct) error {
	stmt, err := models.DB.Prepare("INSERT INTO topics(category_id,user_id,title,excerpt,slug,body,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(m.CategoryId, m.UserId, m.Title, m.Excerpt, m.Slug, m.Body, m.CreateTime, m.UpdatedAt)
	return err
}
