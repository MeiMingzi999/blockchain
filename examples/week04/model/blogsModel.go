package model

//创建时使用
//package main

import (
	"fmt"
	"time"

	"week04/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 用户信息
type User struct {
	ID       uint
	Username string
	Password string
	Email    string

	Posts    []Post    `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:UserID"`
}

// 文章信息
type Post struct {
	ID        uint
	Title     string
	Content   string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment `gorm:"foreignKey:PostID"`
}

// 文章评论信息
type Comment struct {
	ID        uint
	Content   string
	UserID    uint
	PostID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	db, err := gorm.Open(sqlite.Open("../db/blogs.db"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	err = db.AutoMigrate(&User{}, &Comment{}, &Post{})
	if err != nil {
		panic("创建表失败: " + err.Error())
	}
	// ========== 种子数据 ==========
	// 1. 创建2个用户
	pwd1, _ := utils.HashPassword("123456")
	pwd2, _ := utils.HashPassword("654321")

	user1 := User{
		Username: "zhangsan",
		Password: pwd1,
		Email:    "zhangsan@test.com",
	}
	user2 := User{
		Username: "lisi",
		Password: pwd2,
		Email:    "lisi@test.com",
	}

	db.Create(&user1)
	db.Create(&user2)
	fmt.Printf("创建用户：%+v, %+v\n", user1, user2)

	// 2. 创建文章（user1发两篇）
	post1 := Post{
		Title:   "Go入门笔记",
		Content: "今天学习了go mod、gorm、gin，写博客项目",
		UserID:  user1.ID,
	}
	post2 := Post{
		Title:   "Gin中间件学习",
		Content: "CORS中间件用来解决浏览器跨域问题",
		UserID:  user1.ID,
	}
	db.Create(&post1)
	db.Create(&post2)
	fmt.Printf("创建文章：%+v, %+v\n", post1, post2)

	// 3. 创建评论
	comment1 := Comment{
		Content: "写的很好，学到了！",
		UserID:  user2.ID,
		PostID:  post1.ID,
	}
	comment2 := Comment{
		Content: "CORS这个坑我之前也踩过",
		UserID:  user2.ID,
		PostID:  post2.ID,
	}
	db.Create(&comment1)
	db.Create(&comment2)
	fmt.Printf("创建评论：%+v, %+v\n", comment1, comment2)

	fmt.Println("初始化数据完成！")

}
