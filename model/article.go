package model

import "time"

//文章信息
type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Title        string    `db:"title"`
	Keywords     string    `db:"keywords"`
	ViewCount    int32     `db:"view_count"`
	CommentCount int32     `db:"comment_count"`
	UserName     string    `db:"username"`
	Summary      string    `db:"summary"`
	CreateTime   time.Time `db:"create_time"`
}

//文章内容
type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
	Category
}

//文章列表
type ArticleRecord struct {
	ArticleInfo
	ArticleDetail
}
