package db

import (
	"blogger/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

//插入分类
func InsertCategory(category *model.Category)(categoryId int64,err error){
	sqlstr := "insert into category(category_name,category_no) values(?,?)"
	result, err := DB.Exec(sqlstr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryId,err = result.LastInsertId()
	return
}
//获取分类列表
func GetCategoryList(categoryIds []int64)(categoryList []*model.Category,err error){
	sqlstr,args,err := sqlx.In("select id,category_name,category_no from category where id in(?)",categoryIds)
	if err != nil{
		return
	}
	err = DB.Select(&categoryList,sqlstr,args...)
	return
}


//获取全部分类列表
func GetAllCategoryList()(categoryList []*model.Category,err error){
	sqlstr := "select id,category_name,category_no from category order by category_no asc"
	err = DB.Select(&categoryList,sqlstr)
	if err != nil {
		return
	}
	return
}
//根据分类ID获取分类列表
func GetCategoryByid(id int64)(category *model.Category,err error){
	category = &model.Category{}
	sqlstr := "select id,category_name,category_no from category where id=?"
	err = DB.Get(category,sqlstr,id)
	if err != nil {
		return
	}
	return
}
//获取相关文章
func GetRelevantArticleByid(articleId int64)(RelevantArticle []*model.RelevantArticle,err error){
	var categoryId int64
	sqlstr := `select category_id from article where id =?`
	err = DB.Get(&categoryId,sqlstr,articleId)
	if err != nil{
		return
	}

	sqlstr = `select id,title from article where category_id = ? and id !=? limit 10 `
	err = DB.Select(&RelevantArticle,sqlstr,categoryId,articleId)
	return
}
//获取分类文章
func GetCategoryArticle(categoryId int64,pageNum,pageSize int )(articleList []*model.ArticleInfo, err error){
	if pageNum <0 || pageSize < 0 {
		err = fmt.Errorf("invalid paramter ,page_num:%d,page_size:%d",pageNum,pageSize)
		return
	}
	sqlstr := `select 
				id,summary,title,view_count,create_time,comment_count,username,category_id 
				from article where status = 1 and category_id = ?				
				order by create_time desc 
				limit ?,? `

	err = DB.Select(&articleList,sqlstr,categoryId,pageNum,pageSize)

	return
}