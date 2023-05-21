package handler

import (
	"class08/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handlers) CreateTodo(c *gin.Context) {
	var todo model.Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.UserID = c.GetInt("userID")
	h.DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

func (h *Handlers) GetTodos(c *gin.Context) {
	var todos []model.Todo
	h.DB.Where("user_id = ?", c.GetInt("userID")).Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func (h *Handlers) UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedTodo model.Todo

	if err := c.BindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	h.DB.Model(&model.Todo{}).Where("id = ?", id).Updates(&updatedTodo)
	c.JSON(http.StatusOK, updatedTodo)
}

func (h *Handlers) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo model.Todo
	h.DB.Where("id = ? AND user_id = ?", id, c.GetInt("userID")).First(&todo)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	h.DB.Where("id = ?", id).Delete(&model.Todo{})
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
