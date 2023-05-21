package handler

import (
	"class09/database"
	"class09/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type MyTodo struct {
	Todo *model.Todo
}

func (m MyTodo) List(c *gin.Context) {
	var todos []model.Todo

	db, err := database.InitDB()
	if err != nil {
		log.Println(err)
	}
	db.Where("user_id = ?", c.GetInt("userID")).Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func (m MyTodo) Create(c *gin.Context) {
	var todo model.Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.UserID = c.GetInt("userID")

	db, err := database.InitDB()
	if err != nil {
		log.Println(err)
	}
	db.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

func (m MyTodo) Update(c *gin.Context) {
	var updatedTodo model.Todo

	if err := c.BindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	db, err := database.InitDB()
	if err != nil {
		log.Println(err)
	}
	db.Model(&model.Todo{}).Where("id = ?", updatedTodo.ID).Updates(&updatedTodo)
	c.JSON(http.StatusOK, updatedTodo)
}

func (m MyTodo) Delete(c *gin.Context) {

	var todo model.Todo
	db, err := database.InitDB()
	if err != nil {
		log.Println(err)
	}

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	db.Where("id = ? AND user_id = ?", todo.ID, c.GetInt("userID")).First(&todo)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	db.Where("id = ?", todo.ID).Delete(&model.Todo{})
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}

