package main

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type user struct {
	ID         uint
	Name       string
	Post       []post `gorm:"foreignKey:UserId"`
	PostCount  uint
	UpdateTime time.Time
	CreateTime time.Time
}

type post struct {
	ID          uint
	UserId      uint
	Context     string
	Comment     []comment `gorm:"foreignKey:PostId"`
	ComentCount uint
	Status      string
	UpdateTime  time.Time
	CreateTime  time.Time
}

type comment struct {
	ID         uint
	PostId     uint
	Text       string
	UpdateTime time.Time
	CreateTime time.Time
}

type commentCount struct {
	PostId  uint
	Context string
	Count   uint
}

func findInfoByUserID(userId uint) {
	db, error := gorm.Open(sqlite.Open("db/test.db"), &gorm.Config{})
	if error != nil {
		fmt.Printf("数据库连接失败!")
	}
	var u user
	result := db.Preload("Post").Preload("Post.Comment").Where("id = ?", userId).First(&u)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("没有找到该用户")
		} else {
			fmt.Printf("查询异常: %v\n", result.Error)
		}
		return
	}

	fmt.Printf("查询结果为: %+v\n", u)
}

// 查询评论最多的文章内容
func findMaxComment() {
	db, error := gorm.Open(sqlite.Open("db/test.db"), &gorm.Config{})
	if error != nil {
		fmt.Printf("数据库连接失败!")
	}
	var commentcounts []commentCount
	db.Model(&comment{}).Select("comments.post_id, posts.context ,COUNT(*) as count").Joins("left join posts on posts.id = comments.post_id").Group("comments.post_id,posts.context").Order("count desc").Scan(&commentcounts)
	fmt.Printf("查询结果为: %+v\n", commentcounts)
	if len(commentcounts) > 0 {
		fmt.Printf("评论最多的文章ID: %v, 内容为: %v\n", commentcounts[0].PostId, commentcounts[0].Context)
	}

}

// 文章创建时,更新用户的文章数量统计
func (post *post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&user{}).
		Where("users.id = ?", post.UserId).
		Update("post_count", gorm.Expr("COALESCE(post_count, 0) + ?", 1)).Error

}

// 删除评论后,更新文章的评论数,如果评论数为0, 那么更新文章的状态为 无评论
func (c *comment) AfterDelete(tx *gorm.DB) error {
	error := tx.Model(&post{}).Where("id = ?", c.PostId).
		Update("coment_count", gorm.Expr("CASE WHEN COALESCE(coment_count,0) - ? > 0 THEN COALESCE(coment_count,0) - ? ELSE 0 END", 1, 1)).
		Error
	if error != nil {
		fmt.Printf("更新失败, post ID: %v\n", c.PostId)
	}
	return tx.Model(&post{}).Where("id = ? and coment_count = 0", c.PostId).Update("status", "无评论").Error
}

func main() {
	db, error := gorm.Open(sqlite.Open("db/test.db"), &gorm.Config{})
	if error != nil {
		fmt.Printf("数据库连接失败!")
	}
	//findInfoByUserID(1)
	//findMaxComment()
	db.AutoMigrate(&user{}, &post{}, &comment{})
	// newPost := post{
	// 	UserId:  1,
	// 	Context: "测试钩子新增帖子",
	// }
	// error = db.Create(&newPost).Error
	// if error != nil {
	// 	fmt.Println("新建帖子失败：", error)
	// } else {
	// 	fmt.Println("新建帖子成功，自动更新用户帖子计数")
	// }

	c1 := comment{PostId: 1, Text: "不错！"}
	db.Create(&c1)
	db.Delete(&c1)

}
