package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin框架的web demo
func main() {

	// 获取路由对象
	router := gin.Default()
	// 发送get请求
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	// 返回JSON请求
	router.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    User{Id: 1001, Name: "张三", Age: 32},
		})
	})

	// 运行
	err := router.Run(":8000")
	if err != nil {
		return
	}

}

// User 用户结构体
type User struct {
	Id   int
	Name string
	Age  int
}
