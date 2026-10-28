package service

import (
	"net/http"
	"week04/utils"

	"week04/model"

	"github.com/gin-gonic/gin"
)

// 用户登陆
func Login(c *gin.Context) {
	var loginReq struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if error := c.ShouldBindJSON(&loginReq); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	db, err := utils.GetDBClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败" + err.Error()})
		return
	}

	//查询用户,校验账号密码
	userInfo := model.User{}
	err = db.Model(&model.User{}).Where("users.username = ?", loginReq.Username).First(&userInfo).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名不存在"})
		return
	}
	checkResult := utils.CheckPassword(loginReq.Password, userInfo.Password)

	if !checkResult {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}

	//生成token
	tokenString, err := utils.GenerateToken(userInfo.ID, userInfo.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "生成token失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":       userInfo.ID,
			"username": userInfo.Username,
		},
	})
	return

}

// 注册
func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "解析参数失败" + err.Error(),
		})
	}

	db, err := utils.GetDBClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接失败" + err.Error()})
		return
	}
	//校验用户名重复
	var existUser model.User
	err = db.Model(&model.User{}).Where("users.username = ?", user.Username).First(&existUser).Error
	if err == nil {
		// 查到用户，用户名重复
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
		return
	}

	//密码加密
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败" + err.Error()})
		return
	}
	user.Password = hashPassword

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据入库失败" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"userId": user.ID})
	return

}
