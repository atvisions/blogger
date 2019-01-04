package logic

import (
	"blogger/dal/db"
	"blogger/model"
	"fmt"
	"math"
)

func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
LABEL:
	for _, article := range articleInfoList {
		categoryId := article.CategoryId
		for _, id := range ids {
			if id == categoryId {
				continue LABEL
			}
		}
		ids = append(ids, categoryId)
	}
	return
}

func GetArticleArticleRecord(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {

	//从数据库中获取文章列表
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)

	if err != nil {
		fmt.Printf("get articlelist failed:%d", err)
		return
	}

	if len(articleInfoList) == 0 {
		return
	}
	//从数据库中获取分类列表
	categoryIds := getCategoryIds(articleInfoList)

	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("get categoryList failed:%d", err)
		return
	}
	//聚合数据

	for _,article := range articleInfoList{
		articleRecord := &model.ArticleRecord{
			ArticleInfo:*article,
		}
		categoryId := article.CategoryId
		for _,category := range categoryList{
			if categoryId == category.Id {
				articleRecord.Category = *category
				break
			}

		}
		articleRecordList = append(articleRecordList,articleRecord)
	}
	fmt.Println(articleRecordList)
	return
}
//插入文章
func InsertArticle (title, content, username,keywords string,categoryId int64)(err error){
	//1、构造articleDetail
	articleDetail := &model.ArticleDetail{}
	articleDetail.Title = title
	articleDetail.Content = content
	articleDetail.UserName = username
	articleDetail.Keywords = keywords
	articleDetail.ArticleInfo.CategoryId = categoryId
	//摘要
	contentUtf8 := []rune(content)
	minLength := int(math.Min(float64(len(contentUtf8)),128.0))
	articleDetail.Summary = string([]rune(content[:minLength]))

	id,err := db.InsertArticle(articleDetail)
	fmt.Println(id)
	return

}
//获取文章内容
func ArticleDetail(articleId int64)(ArticleDetail *model.ArticleDetail){
	//获取文章信息
	ArticleDetail,err := db.GetArticleDetail(articleId)
	if err != nil {
		fmt.Printf("get ArticleDetail failed:%d", err)
		return
	}
	//获取文章对应的分类
	category, err := db.GetCategoryByid(ArticleDetail.ArticleInfo.CategoryId)
	if err != nil {
		fmt.Printf("get ArticleDetail failed:%d", err)
		return
	}
	ArticleDetail.Category = *category
	return

}


