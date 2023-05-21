package posts

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//记录日志
		fmt.Println("start")

		//业务处理
		c.Next()
	}
}

func Router(r *gin.Engine) {
	posts := r.Group("/posts")

	posts.Use(Logger())

	posts.GET("/", func(c *gin.Context) {
		//返回所有文章信息
	})

	posts.POST("/", func(c *gin.Context) {
		//创建新文章
	})

	posts.PUT("/:id", func(c *gin.Context) {
		//更新指定文章信息
	})

	posts.DELETE("/:id", func(c *gin.Context) {
		//删除指定文章信息
	})
}
