package logic

import (
	"blogger/dal/db"
	"blogger/model"
	"fmt"
)

func GetAllCategoryList()(categoryList []*model.Category,err error){
	categoryList,err = db.GetAllCategoryList()
	if err != nil {
		return
	}
	return
}

//获取相关文章
func RelevantArticle(articleId int64)(RelevantArticle []*model.RelevantArticle,err error){
	RelevantArticle ,err = db.GetRelevantArticleByid(articleId)
	if err != nil{
		return
	}

	return

}
//获取分类文章
func CategoryArticle(categoryId int64,pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {

	//从数据库中获取文章列表
	articleInfoList, err := db.GetCategoryArticle(categoryId,pageNum, pageSize)

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
