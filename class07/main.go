package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

const secretKey = "abc123"

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"  gorm:"unique"`
	Password string `json:"password"`
}

type Todo struct {
	gorm.Model
	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status string `json:"status"`
	UserID int    `json:"user_id"`
}

func initDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{}, &Todo{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func signUP(c *gin.Context, db *gorm.DB) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser User
	db.Where("username = ?", user.Username).First(&existingUser)
	if existingUser.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = string(hashedPassword)
	db.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func signIn(c *gin.Context, db *gorm.DB) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser User
	db.Where("username = ?", user.Username).First(&existingUser)
	if existingUser.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       existingUser.ID,
		"username": existingUser.Username,
	})

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}

func authenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("userID", int(claims["id"].(float64)))
		c.Next()
	}
}

func createTodo(c *gin.Context, db *gorm.DB) {
	var todo Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.UserID = c.GetInt("userID")
	db.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

func getTodos(c *gin.Context, db *gorm.DB) {
	var todos []Todo
	db.Where("user_id = ?", c.GetInt("userID")).Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func updateTodo(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedTodo Todo

	if err := c.BindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	db.Model(&Todo{}).Where("id = ?", id).Updates(&updatedTodo)
	c.JSON(http.StatusOK, updatedTodo)
}

func deleteTodo(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo Todo
	db.Where("id = ? AND user_id = ?", id, c.GetInt("userID")).First(&todo)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	db.Where("id = ?", id).Delete(&Todo{})
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()

	r.POST("/signup", func(c *gin.Context) {
		signUP(c, db)
	})

	r.POST("/signin", func(c *gin.Context) {
		signIn(c, db)
	})

	authorized := r.Group("/")
	authorized.Use(authenticationMiddleware())
	{
		authorized.POST("/todos", func(c *gin.Context) {
			createTodo(c, db)
		})
		authorized.GET("/todos", func(c *gin.Context) {
			getTodos(c, db)
		})
		authorized.PUT("/todos/:id", func(c *gin.Context) {
			updateTodo(c, db)
		})
		authorized.DELETE("/todos/:id", func(c *gin.Context) {
			deleteTodo(c, db)
		})
		authorized.POST("/signout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
		})
	}

	r.Run(":8888")
}
