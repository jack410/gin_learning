package handler

import (
	"class09/database"
	"class09/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type MyCustomer struct {
	Customer *model.Customer
}

func (m *MyCustomer) List(c *gin.Context) {
	var customers []model.Customer

	db, err := database.InitDB()
	if err != nil {
		log.Println(err)
	}
	db.Where("user_id = ?", c.GetInt("userID")).Find(&customers)
	c.JSON(http.StatusOK, customers)
}

func (m *MyCustomer) Create(c *gin.Context) {
	var customer model.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer.UserID = c.GetInt("userID")

	db, err := database.InitDB()
	if err != nil {
		log.Println(err)
	}
	db.Create(&customer)
	c.JSON(http.StatusCreated, customer)
}

func (m *MyCustomer) Update(c *gin.Context) {
	var updatedCustomer model.Customer

	if err := c.BindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	db, err := database.InitDB()
	if err != nil {
		log.Println(err)
	}
	db.Model(&model.Customer{}).Where("id = ?", updatedCustomer.ID).Updates(&updatedCustomer)
	c.JSON(http.StatusOK, updatedCustomer)
}

func (m *MyCustomer) Delete(c *gin.Context) {

	var customer model.Customer
	db, err := database.InitDB()
	if err != nil {
		log.Println(err)
	}

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	db.Where("id = ? AND user_id = ?", customer.ID, c.GetInt("userID")).First(&customer)
	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	db.Where("id = ?", customer.ID).Delete(&model.Customer{})
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
