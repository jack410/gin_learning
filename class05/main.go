package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//创建一级路由分组
	v1 := r.Group("/v1")

	//创建二级路由分组
	users := v1.Group("/users")

	//定义二级路由分组里的路由
	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "返回所有用户信息",
		})
	})

	users.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "创建新用户",
		})
	})

	users.PUT("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "更新指定用户",
		})
	})

	users.DELETE("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "删除指定用户",
		})
	})

	//创建二级路由分组
	posts := v1.Group("/posts")

	posts.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "返回所有文章信息",
		})
	})

	posts.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "创建新文章",
		})
	})

	posts.PUT("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "更新指定文章",
		})
	})

	posts.DELETE("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "删除指定文章",
		})
	})
	
	r.Run(":8888")

}
