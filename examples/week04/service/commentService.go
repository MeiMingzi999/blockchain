package service

import (
	"fmt"
	"net/http"
	"week04/model"
	"week04/utils"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var reqParam struct {
		PostId  uint   `json:"postId" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&reqParam); err != nil {
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

	db, err := utils.GetDBClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败" + err.Error()})
		return
	}
	comment := model.Comment{
		UserID:  loginUserId,
		PostID:  reqParam.PostId,
		Content: reqParam.Content,
	}
	if err = db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "评论插入失败" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "评论成功",
	})

}

// 某篇文章的所有评论
func GetCommentList(c *gin.Context) {
	db, err := utils.GetDBClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败" + err.Error()})
		return
	}
	postId := c.Param("id")

	var commentList []model.Comment
	err = db.Model(&model.Comment{}).Where("post_id = ?", postId).Scan(&commentList).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询评论失败" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"data": commentList,
	})

}
