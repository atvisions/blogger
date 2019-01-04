package main

import (
	"blogger/controller"
	"blogger/dal/db"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//打开数据库
	db.InitDB()
	//
	router := gin.Default()
	router.Static("/static","./static")
	router.LoadHTMLGlob("views/*")
	//首页路由
	router.GET("/",controller.IndexHandle)
	//文章发布页面路由
	router.GET("/post",controller.PostArticle)
	//文章发布路由
	router.POST("/post/submit",controller.ArticleSubmit)
	//文章发布路由
	router.GET("/article/detail",controller.ArticleDetail)
	//文章分类列表
	router.GET("/category/category",controller.CategoryList)
	//更新网站
	router.GET("/webinfo",controller.WebInfo)
	//更新网站
	//router.POST("/webinfo/post",controller.UpWebInfo)
	// 绑定端口是8888
	router.Run(":8888")

}
