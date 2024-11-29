# Gin 框架

------

***创建时间 2024-11-24***

## 文件结构

```bash
├── assets # 静态资源
├── docs # 文档
├── gin # gin web framework 框架
│   ├── controllers # 控制器
│   ├── main.go # gin 入口文件
│   ├── models # SQL 
│   ├── routers # 路由
├── go.sum # 三方库
├── grammar # 基础语法写法
├── main.go # 基础语法入口

```





## 热更新

```go
go install github.com/air-verse/air@latest
// 启动命令
air
```



## 基础路由引擎

```go
import("github.com/gin-gonic/gin")
func main () {
  // 创建一个默认的路由引擎
	r := gin.Default()
  // 配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值：%v", "你好golang gin")
	})
  // 路由分组
  routers.InitRouters(r)
  // 启动一个web服务
	runErr := r.Run(":9981")
	if runErr != nil {
		fmt.Println("gin 启动失败！！！")
	}
}
```

### 路由分组

```go
// routers/initRouters
package routers
func InitRouters(r *gin.Engine) {
	BasicRoutersInit(r)
	AdminRoutersInit(r)
	ApiRoutersInit(r)
}

// routers/basicRouters
package routers
func BasicRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
  // 中间件
  defaultRouters.Use(middleware.MiddleW)
	{
		// 配置路由，单独路由中间件
		defaultRouters.GET("/", middleware.HomeMiddle, func(c *gin.Context) {
			c.String(200, "值：%v", "你好golang gin")
		})

		defaultRouters.GET("/ping", func(c *gin.Context) {
			// http.StatusOK 200 成功状态码
			c.JSON(http.StatusOK, gin.H{
				"message": "pong -> air-verse/air 启动",
			})
		})
	}
}
```

## 中间件

1. Next()：表示跳过当前中间件剩余内容，去执行下一个中间件
2. return：终止执行当前中间件剩余内容，执行下一个中间件
3. Abort()：只执行当前中间件，操作完成后，以出栈的顺序，一次返回返回上一级中间件

```go
// middleware/middleW.go
package middleware

// 默认中间件写法
func MiddleW(c *gin.Context) {
	fmt.Println("middleW 中间件函数")
	// 表示跳过当前中间件剩余内容，去执行下一个中间件
	//c.Next()
}

/*
中间件 第二种写法。需要执行
*/
func MiddleW2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleW2 中间函数")
		//c.Next()
	}
}

/*
单独路由中间件
*/
func HomeMiddle(c *gin.Context) {
	fmt.Println("特定路由的中间件 ---> home")
	c.Next()
}
```







# SQL

------

***创建时间：2024-11-24***

```go
  "gorm.io/driver/mysql"
  "gorm.io/gorm"

```









