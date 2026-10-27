package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

func main() {
	findInfoByUserID(1)
}
