package service

import (
	"fmt"
	"net/http"
	"time"
	"week04/model"
	"week04/utils"

	"github.com/gin-gonic/gin"
)

// 创建文章
func CreatePost(c *gin.Context) {
	var post model.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章参数解析失败" + err.Error()})
		return
	}

	db, err := utils.GetDBClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败" + err.Error()})
		return
	}
	userIdVal, ok := c.Get("userId")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户信息不存在"})
		return
	}
	fmt.Printf("userIdVal: %v\n", userIdVal)

	userId, ok := userIdVal.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户ID格式错误"})
		return
	}
	post.UserID = userId
	db.Create(&post)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "发布文章成功",
		"post": post,
	})

}

// 获取所有文章列表
func GetPostList(c *gin.Context) {
	db, err := utils.GetDBClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败" + err.Error()})
		return
	}
	userIdVal, ok := c.Get("userId")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户信息不存在"})
		return
	}
	fmt.Printf("userIdVal: %v\n", userIdVal)

	userId, ok := userIdVal.(uint)

	var postList []struct {
		ID    uint
		Title string
	}
	err = db.Model(&model.Post{}).Where("posts.user_id = ?", userId).Scan(&postList).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询文章失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": postList})
}

// 查询单个文章详细信息
func GetPostInfo(c *gin.Context) {
	db, err := utils.GetDBClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败" + err.Error()})
		return
	}
	id := c.Param("id")
	var postInfo struct {
		ID        uint
		Title     string
		Content   string
		UserID    uint
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	err = db.Model(&model.Post{}).Where("id =?", id).Scan(&postInfo).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询文章失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": postInfo})

}

// 修改文章
func UpdatePost(c *gin.Context) {
	db, err := utils.GetDBClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败" + err.Error()})
		return
	}
	// 接收前端参数
	var req struct {
		ID      uint   `json:"id" binding:"required"`
		Title   string `json:"title"`
		Content string `json:"content"`
		UserID  int    `json:"userId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数解析失败: " + err.Error()})
		return
	}

	// 获取当前登录用户
	userIdVal, ok := c.Get("userId")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户信息不存在"})
		return
	}
	fmt.Printf("userIdVal: %v\n", userIdVal)

	loginUserId, ok := userIdVal.(uint)

	// 先查这篇文章，校验是否存在 + 是否是本人的文章
	var post model.Post
	err = db.First(&post, req.ID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在"})
		return
	}
	if post.UserID != loginUserId {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权修改他人文章"})
		return
	}

	updateData := map[string]any{}
	if req.Title != "" {
		updateData["title"] = req.Title
	}
	if req.Content != "" {
		updateData["content"] = req.Content
	}
	if len(updateData) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "没有需要更新的内容"})
		return
	}

	err = db.Model(&post).Updates(updateData).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "修改文章成功",
		"post": post,
	})

}

// 删除
func DeletePost(c *gin.Context) {
	db, err := utils.GetDBClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败" + err.Error()})
		return
	}
	id := c.Param("id")

	// 获取当前登录用户
	userIdVal, ok := c.Get("userId")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户信息不存在"})
		return
	}
	fmt.Printf("userIdVal: %v\n", userIdVal)

	loginUserId, ok := userIdVal.(uint)

	// 先查这篇文章，校验是否存在 + 是否是本人的文章
	var post model.Post
	err = db.First(&post, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在"})
		return
	}
	if post.UserID != loginUserId {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权删除他人文章"})
		return
	}
	err = db.Delete(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除文章失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "删除文章成功",
	})

}
