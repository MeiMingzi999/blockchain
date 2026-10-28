# week04 博客项目接口文档

## 项目简介

基于 Gin + GORM + SQLite + JWT 实现的简易博客后端，包含用户注册登录、文章增删改查、评论功能，接口统一携带 JWT 鉴权。

> 基础信息
>
> - 服务地址：`http://127.0.0.1:8080`
> - Content-Type：`application/json`
> - 鉴权方式：请求头 `Authorization: Bearer {token}`
> - 所有 `/api/post`、`/api/comment` 分组下接口**必须携带 token**

## 接口列表

### 一、用户模块（无需登录）

#### 1. 用户登录

- 请求方式：`POST`
- 接口地址：`/api/login`
- 请求体

```json
{
    "username":"zhangsan",
    "password":"123456"
}
```

- 成功响应

```json
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTMyNjA3MjQsImlkIjoxLCJ1c2VybmFtZSI6InpoYW5nc2FuIn0.VWjyCSsrdxV5wjoBpsV6dS4UaChCPpD6FB0ZNVlO9-s",
	"user": {
		"id": 1,
		"username": "zhangsan"
	}
}
```

![image-20261028185626547](./static/image-20261028185626547.png)

#### 2. 用户注册

- 请求方式：`POST`
- 接口地址：`/api/register`
- 请求体

```json
{
    "username":"zhangsan11",
    "password":"123456",
    "email":"123@qq.com"
}
```

- 成功响应

```json
{
	"userId": 4
}
```

------

![image-20261028185732099](./static/image-20261028185732099.png)

### 二、文章模块（需要登录，Header 携带 Token）

> 前缀：`/api/post`

#### 1. 创建文章

- 请求方式：`POST`
- 接口地址：`/api/post/createPost`
- 请求体

```json
{
    "title":"测试新建",
    "content":"hahahahahaha"
}
```

- 成功响应

```json
{
	"msg": "发布文章成功",
	"post": {
		"ID": 3,
		"Title": "测试新建",
		"Content": "hahahahahaha",
		"UserID": 1,
		"CreatedAt": "2026-10-28T17:25:12.795555+08:00",
		"UpdatedAt": "2026-10-28T17:25:12.795555+08:00",
		"Comments": null
	}
}
```

![image-20261028185909541](./static/image-20261028185909541.png)

#### 2. 获取当前用户文章列表

- 请求方式：`GET`
- 接口地址：`/api/post/getPostList`
- 成功响应

```json
{
	"data": [
		{
			"ID": 1,
			"Title": "Go入门笔记"
		},
		{
			"ID": 2,
			"Title": "Gin中间件学习"
		},
		{
			"ID": 3,
			"Title": "测试新建"
		}
	]
}
```

![image-20261028185947452](./static/image-20261028185947452.png)

#### 3. 获取单篇文章详情

- 请求方式：`GET`
- 接口地址：`/api/post/getPostInfo/:id`
- 路径参数：`id` 文章 ID (文章ID: 1)
- 成功响应

```json
{
	"data": {
		"ID": 1,
		"Title": "Go入门笔记",
		"Content": "今天学习了go mod、gorm、gin，写博客项目",
		"UserID": 1,
		"CreatedAt": "2026-10-28T02:17:12.929089+08:00",
		"UpdatedAt": "2026-10-28T02:17:12.929089+08:00"
	}
}
```

![image-20261028190046944](./static/image-20261028190046944.png)

#### 4. 修改文章

- 请求方式：`POST`
- 接口地址：`/api/post/updatePost`
- 请求体

```json
{
		"ID": 1,
		"Title": "Go入门笔记(修改版)",
		"Content": "今天学习了go mod、gorm、gin，写博客项目"
}
```

> 权限校验：只能修改自己发布的文章

- 成功响应

```json
{
	"msg": "修改文章成功",
	"post": {
		"ID": 1,
		"Title": "Go入门笔记(修改版)",
		"Content": "今天学习了go mod、gorm、gin，写博客项目",
		"UserID": 1,
		"CreatedAt": "2026-10-28T02:17:12.929089+08:00",
		"UpdatedAt": "2026-10-28T18:09:05.440404+08:00",
		"Comments": null
	}
}
```

![image-20261028190144148](./static/image-20261028190144148.png)

#### 5. 删除文章

- 请求方式：`DELETE`
- 接口地址：`/api/post/deletePost/:id`
- 路径参数：`id` 文章 ID  (删除文章id : 3)

> 权限校验：只能删除自己发布的文章

- 成功响应

```json
{
    "msg": "删除文章成功"
}
```

------

![image-20261028190238210](./static/image-20261028190238210.png)

### 三、评论模块（需要登录，Header 携带 Token）

> 前缀：`/api/comment`

#### 1. 创建评论

- 请求方式：`POST`
- 接口地址：`/api/comment/createComment`
- 请求体示例

```json
{
		"postId": 1,
		"content": "很棒棒"
}
```

- 成功响应

```json
{
    "msg": "评论成功"
}
```

![image-20261028190322970](./static/image-20261028190322970.png)

#### 2. 获取文章评论列表

- 请求方式：`GET`
- 接口地址：`/api/comment/getCommentList/:id`
- 路径参数：`id` 文章 ID
- 成功响应

```json
{
	"data": [
		{
			"ID": 1,
			"Content": "写的很好，学到了！",
			"UserID": 2,
			"PostID": 1,
			"CreatedAt": "2026-10-28T02:17:12.929824+08:00",
			"UpdatedAt": "2026-10-28T02:17:12.929824+08:00"
		},
		{
			"ID": 3,
			"Content": "很棒棒",
			"UserID": 1,
			"PostID": 1,
			"CreatedAt": "2026-10-28T18:40:37.027184+08:00",
			"UpdatedAt": "2026-10-28T18:40:37.027184+08:00"
		}
	],
	"msg": "查询成功"
}
```

------

![image-20261028190429999](./static/image-20261028190429999.png)

## 四、项目启动说明

### 1. 安装依赖

```
go mod tidy
```

### 2. 启动项目

```
go run main.go
```

服务默认监听 `:8080`

