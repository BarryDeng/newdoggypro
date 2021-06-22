package main

import (
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type (
	userModel struct {
		gorm.Model
		Name     string `json:"name"`
		Password int    `json:"password"`
	}

	transformedUser struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}
)

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db.AutoMigrate(&userModel{})
}

func main() {
	router := gin.Default()
	router.Static("/", "./web/dist")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	v1Users := router.Group("/api/v1/users")
	{
		v1Users.POST("/", createUsers)
		v1Users.GET("/", fetchAllUsers)
		v1Users.GET("/:id", fetchUser)
		v1Users.PUT("/:id", updateUsers)
		v1Users.DELETE("/:id", deletaUser)
	}
	router.Run(":8000")
}
