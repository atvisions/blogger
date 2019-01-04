package db

import (
	"blogger/model"
	"fmt"
)
//获取文章
func ArticleQuery(){
	article := make([]model.ArticleDetail, 0)
	sqlStr := "select id,category_id,title,keywords,content,view_count,comment_count,username,status,summary,create_time from article where id>?"
	err := DB.Select(&article, sqlStr, 0)
	if err != nil {
		fmt.Println(err)
	}
	for i,v :=range article {
		fmt.Println(i,v)
	}
}
//插入文章,要选择完整的结构体

func InsertArticle(article *model.ArticleDetail)(articleId int64, err error){
	if article == nil{
		err = fmt.Errorf("文章内容不能为空")
		return
	}
	sqlstr := `insert into 
				article(title,content,username,category_id,summary,view_count,comment_count,keywords) 
				values(?,?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlstr,article.Title,article.Content,article.UserName, article.CategoryId, article.Summary, article.ViewCount,article.CommentCount,article.Keywords)
	if err != nil{
		return
	}
	//返回最后插入的ID
	articleId ,err = result.LastInsertId()

	return
}
//获取文章列表
func GetArticleList(pageNum,pageSize int )(articleList []*model.ArticleInfo, err error){
	if pageNum <0 || pageSize < 0 {
		err = fmt.Errorf("invalid paramter ,page_num:%d,page_size:%d",pageNum,pageSize)
		return
	}
	sqlstr := `select 
				id,summary,title,keywords,view_count,create_time,comment_count,username,category_id
				from article where status = 1 				
				order by create_time desc 
				limit ?,? `

	err = DB.Select(&articleList,sqlstr,pageNum,pageSize)

	return
}
//获取文章内容
func GetArticleDetail(articleId int64)(articleDetail *model.ArticleDetail, err error){
	if articleId < 0 {
		err = fmt.Errorf("invalid paramter ,articleId:%d",articleId)
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlstr := `select 
				id,summary,title,keywords,view_count,create_time,comment_count,username,category_id,content 
				from article where status = 1 and id =?`

	err = DB.Get(articleDetail,sqlstr,articleId)

	return
}
