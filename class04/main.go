package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	//创建路由分组
	api := r.Group("/api")

	//添加中间件
	api.Use(Logger())

	//定义路由分组里面的路由
	//GET users
	api.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "返回所以用户信息",
		})
	})

	//POST users
	api.POST("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "创建新用户",
		})
	})

	//PUT /users/:id
	api.PUT("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "更新指定用户信息",
		})
	})

	//DELETE /users/:id
	api.DELETE("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "删除指定用户",
		})
	})

	r.Run(":8888")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		//请求前的日志
		log.Printf("[%s] %s %s \n", t.Format("2006-01-02 15:04:05"), c.Request.Method, c.Request.URL.Path)

		//处理请求
		c.Next()

		//请求后记录响应日志
		log.Printf("[%s] %s %s %s\n", t.Format("2006-01-02 15:04:05"), c.Request.Method, c.Request.URL.Path, time.Since(t))
	}
}
