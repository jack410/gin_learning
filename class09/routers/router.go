package routers

import (
	"class09/database"
	"class09/handler"
	"class09/middleware"
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

	resources := r.Group("/")
	resources.Use(middleware.Auth())
	resources.Any("/:resourcetype", restHandler)
	resources.POST("/signout", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
	})

	return r
}

func restHandler(c *gin.Context) {
	m := map[string]handler.ResourceMethod{
		"todo":     &handler.MyTodo{},
		"customer": &handler.MyCustomer{},
	}
	resourcetype := c.Param("resourcetype")

	switch c.Request.Method {
	case "GET":
		m[resourcetype].List(c)
	case "POST":
		m[resourcetype].Create(c)
	case "PUT":
		m[resourcetype].Update(c)
	case "DELETE":
		m[resourcetype].Delete(c)
	}
}

func NewServer() Server {
	return setupRouter()
}
