package users

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
	users := r.Group("/users")

	users.Use(Logger())

	users.GET("/", func(c *gin.Context) {
		//返回所有用户信息
	})

	users.POST("/", func(c *gin.Context) {
		//创建新用户
	})

	users.PUT("/:id", func(c *gin.Context) {
		//更新指定用户信息
	})

	users.DELETE("/:id", func(c *gin.Context) {
		//删除指定用户信息
	})
}
