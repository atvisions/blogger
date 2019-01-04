package model

type RelevantArticle struct {
	Id    int64  `db:"id"`
	Title string `db:"title"`
}
