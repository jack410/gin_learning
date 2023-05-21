package routers

import (
	"class08/database"
	"class08/handler"
	"class08/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server interface {
	Run(addr ...string) (err error)
}

func setupRouter() *gin.Engine {
	db, err := database.InitDB()
	if err != nil {
		log.Println(err)
	}

	h := handler.Handlers{DB: db}

	r := gin.Default()

	r.POST("/signup", h.SignUp)

	r.POST("/signin", h.SignIn)

	authorized := r.Group("/")
	authorized.Use(middleware.Auth())

	{
		authorized.POST("/todos", h.CreateTodo)
		authorized.GET("/todos", h.GetTodos)
		authorized.PUT("/todos/:id", h.UpdateTodo)
		authorized.DELETE("/todos/:id", h.DeleteTodo)
		authorized.POST("/signout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
		})
	}

	return r
}

func NewServer() Server {
	return setupRouter()
}
