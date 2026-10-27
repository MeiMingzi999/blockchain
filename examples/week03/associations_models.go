package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type user struct {
	ID         uint
	Name       string
	Post       []post `gorm:"foreignKey:UserId"`
	UpdateTime time.Time
	CreateTime time.Time
}

type post struct {
	ID         uint
	UserId     uint
	Context    string
	Comment    []comment `gorm:"foreignKey:PostId"`
	UpdateTime time.Time
	CreateTime time.Time
}

type comment struct {
	ID         uint
	PostId     uint
	Text       string
	UpdateTime time.Time
	CreateTime time.Time
}

func main() {
	// 打开sqlite数据库，文件名叫test.db
	db, err := gorm.Open(sqlite.Open("db/test.db"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	// 自动迁移：不存在表则创建；存在只新增字段，不删改原有数据
	err = db.AutoMigrate(&user{}, &post{}, &comment{})
	if err != nil {
		panic("迁移建表失败: " + err.Error())
	}

	fmt.Println("表创建完成！")

	//插入数据
	users := []user{
		{Name: "张三", CreateTime: time.Now(), UpdateTime: time.Now()},
		{Name: "李四", CreateTime: time.Now(), UpdateTime: time.Now()},
		{Name: "王五", CreateTime: time.Now(), UpdateTime: time.Now()},
	}
	if err = db.Create(users).Error; err == nil {
		fmt.Printf("插入表失败")
	}
	// 插入多条帖子，UserId 关联用户ID
	posts := []post{
		{UserId: 1, Context: "今天学习GORM关联查询", CreateTime: time.Now(), UpdateTime: time.Now()},
		{UserId: 1, Context: "GORM一对多关系挺有意思", CreateTime: time.Now(), UpdateTime: time.Now()},
		{UserId: 2, Context: "SQLite轻量好用", CreateTime: time.Now(), UpdateTime: time.Now()},
		{UserId: 3, Context: "Go语言开发效率很高", CreateTime: time.Now(), UpdateTime: time.Now()},
	}
	if err = db.Create(&posts).Error; err != nil {
		panic("插入帖子失败: " + err.Error())
	}
	fmt.Printf("插入 %d 条帖子成功\n", len(posts))

	// 插入多条评论，PostId 关联帖子ID
	comments := []comment{
		{PostId: 1, Text: "学到了！", CreateTime: time.Now(), UpdateTime: time.Now()},
		{PostId: 1, Text: "我也正在看这个", CreateTime: time.Now(), UpdateTime: time.Now()},
		{PostId: 2, Text: "是的，预加载很方便", CreateTime: time.Now(), UpdateTime: time.Now()},
		{PostId: 3, Text: "不用单独装数据库", CreateTime: time.Now(), UpdateTime: time.Now()},
		{PostId: 4, Text: "同意，写起来简洁", CreateTime: time.Now(), UpdateTime: time.Now()},
	}
	if err = db.Create(&comments).Error; err != nil {
		panic("插入评论失败: " + err.Error())
	}
	fmt.Printf("插入 %d 条评论成功\n", len(comments))

	fmt.Println("所有数据插入完成！")
}
