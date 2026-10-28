package main

import (
	"week04/middleware"
	"week04/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	r.POST("/api/login", service.Login)
	r.POST("/api/register", service.Register)
	//需文章分组, 需要登陆校验
	postApi := r.Group("/api/post")
	postApi.Use(middleware.AuthMiddleware())
	{
		postApi.POST("/createPost", service.CreatePost)
		postApi.GET("/getPostList", service.GetPostList)
		postApi.GET("/getPostInfo/:id", service.GetPostInfo)
		postApi.POST("/updatePost", service.UpdatePost)
		postApi.DELETE("/deletePost/:id", service.DeletePost)
	}

	//评论校验, 需要登陆校验
	commentApi := r.Group("/api/comment")
	commentApi.Use(middleware.AuthMiddleware())
	{
		commentApi.POST("/createComment", service.CreateComment)
		commentApi.GET("/getCommentList/:id", service.GetCommentList)
	}

	r.Run(":8080")

}
