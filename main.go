package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open("root:123456@/demo?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db.AutoMigrate(&userModel{})
	db.AutoMigrate(&pictureModel{})
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	router.Static("/game", "./web/dist")

	v1Users := router.Group("/api/v1/users")
	{
		v1Users.POST("/", createUser)
		v1Users.GET("/", fetchAllUsers)
		v1Users.GET("/:id", fetchUser)
		v1Users.PUT("/:id", updateUser)
		v1Users.DELETE("/:id", deleteUser)
	}

	v1Pictures := router.Group("/api/v1/pictures")
	{
		v1Pictures.POST("/", createPicture)
		v1Pictures.GET("/", fetchAllPictures)
		v1Pictures.GET("/:id", fetchPicture)
		v1Pictures.PUT("/:id", updatePicture)
		v1Pictures.DELETE("/:id", deletePicture)
	}
	router.Run(":8000")
}
