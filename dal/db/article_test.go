package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func init()  {
	InitDB()
}


/*func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{}
	article.Title = "这里是文章的标题22222"
	article.Content = "文章内容"
	article.Summary = "文章摘要"
	article.ArticleInfo.ViewCount = 1
	article.Category.Id = 2
	article.CommentCount = 1
	article.UserName = "caicai"
	article.Keywords = "文章关键字"
	article.ArticleInfo.CreateTime = time.Now()

	articleId,err := InsertArticle(article)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(articleId)
		//InsertArticle(article *model.ArticleDetail)(articleId int64, err error){}

}*/
/*func TestGetArticleList(t *testing.T) {
	articleList,err := GetArticleList(0,15)
	if err != nil {
		fmt.Println(err)
	}

	t.Logf("get article: %d",len(articleList))
	for _,v :=  range articleList{
		t.Logf("%d,%#v",v.Id,v)
	}
}*/

func TestGetArticleDetail(t *testing.T) {
	articleDetail, err := GetArticleDetail(6)
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("%v",articleDetail)

}


