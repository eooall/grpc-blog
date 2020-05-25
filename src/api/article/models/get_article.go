package models

import (
	"fmt"
	pb "grpc_blog/src/api/article/proto"
	"grpc_blog/src/models"
	"time"
)

func GetArticle(page, limit int) ([]*pb.ArticleData, error) {
	stmt, err := models.DB.Prepare(`SELECT a.id,b.name,b.avatar,a.title,a.excerpt,a.slug,a.created_at,count(c.id) as total,IFNULL(d.name,'空') as type
	FROM topics as a
	LEFT JOIN users b
	ON a.user_id = b.id
	LEFT JOIN comments as c
	ON a.id = c.topic_id
	LEFT JOIN categories as d
	ON a.category_id = d.id
	GROUP BY a.id
	ORDER BY a.created_at DESC
	LIMIT ?,?;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	query, err := stmt.Query(page, limit)
	if err != nil {
		return nil, err
	}

	l := make([]*pb.ArticleData, 0)
	for query.Next() {
		temp := pb.ArticleData{}
		var t = time.Time{}
		err = query.Scan(&temp.Id, &temp.Name, &temp.Avatar, &temp.Title, &temp.Excerpt, &temp.Slug, &t, &temp.Total, &temp.ArticleType)
		temp.CreatedAt = t.Format("2006-01-02 15:04:05")
		if err != nil {
			fmt.Println(err)
			continue
		}
		l = append(l, &temp)
	}
	return l, nil
}

func GetTotal() (int64, error) {
	stmt, err := models.DB.Prepare("select count(id) from topics")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	row := stmt.QueryRow()
	var total int64
	row.Scan(&total)
	return total, nil
}
