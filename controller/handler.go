package controller

import (
	"blogger/dal/db"
	"blogger/logic"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//首页
func IndexHandle(c *gin.Context) {
	//文章列表
	articleRecordList, err := logic.GetArticleArticleRecord(0, 15)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	//分类云
	categoryList, err := logic.GetAllCategoryList()
	//网站信息
	webInfo, err := db.GetWebInfo()

	var m map[string]interface{} = make(map[string]interface{}, 0)
	m["articleList"] = articleRecordList
	m["categoryList"] = categoryList
	m["webInfo"] = webInfo

	c.HTML(http.StatusOK, "views/index.html", m)
}

//获取发布文章页面分类ID
func PostArticle(c *gin.Context) {
	categoryList, err := logic.GetAllCategoryList()
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	c.HTML(http.StatusOK, "views/post.html", categoryList)
}

//发布文章
func ArticleSubmit(c *gin.Context) {
	//获取发布页面字段内容
	title := c.PostForm("title")
	username := c.PostForm("username")
	content := c.PostForm("content")
	keywords := c.PostForm("keywords")
	//uploadImage := c.PostForm("/static/upload/UpImages")
	//字符串转INT
	categoryIdStr := c.PostForm("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//插入文章
	err = logic.InsertArticle(title, content, username, keywords, categoryId)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	//发布成功重定向
	c.Redirect(http.StatusMovedPermanently, "/")
}

//文章详情
func ArticleDetail(c *gin.Context) {
	articleIdStr := c.Query("id")
	//字符串转INT
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	ArticleDetail := logic.ArticleDetail(articleId)

	RelevantArticle, err := logic.RelevantArticle(articleId)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}

	var m map[string]interface{} = make(map[string]interface{}, 10)
	m["detail"] = ArticleDetail
	m["relevantArticle"] = RelevantArticle

	c.HTML(http.StatusOK, "views/detail.html", m)
}

//分类文章列表
func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)

	articleInfo, err := logic.CategoryArticle(categoryId, 0, 15)
	if err != nil {
		return
	}
	//分类名称
	categoryName, err := db.GetCategoryByid(categoryId)
	//分类云
	categoryList, err := logic.GetAllCategoryList()
	//网站信息
	webInfo, err := db.GetWebInfo()

	var m map[string]interface{} = make(map[string]interface{}, 0)
	m["articleInfo"] = articleInfo
	m["categoryList"] = categoryList
	m["category"] = categoryName
	m["webInfo"] = webInfo
	c.HTML(http.StatusOK, "views/category.html", m)

}

//网站更新页面
func WebInfo(c *gin.Context) {
	webInfo, err := db.GetWebInfo()
	if err != nil {
		return
	}
	c.HTML(http.StatusOK, "views/webinfo.html", webInfo)
}

/*//网站内容更新
func UpWebInfo(c *gin.Context) {
	web_name := c.PostForm("web_name")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	url := c.PostForm("url")

	webInfo, err := logic.UpWebInfo(web_name, keywords, description, url)
	if err != nil {
		return
	}

}
*/