package main

import (
	"class06/modules/posts"
	"class06/modules/users"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	users.Router(r)

	posts.Router(r)

	r.Run(":9999")
}
