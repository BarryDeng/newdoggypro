package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
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

type (
	user struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		AutoLogin bool   `json:"autoLogin"`
		Type      string `json:"type"`
	}
)

func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Use(static.Serve("/", static.LocalFile("./web/dist", true)))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	router.GET("/api/currentUser", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("User") == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": gin.H{
					"isLogin": false,
				},
				"errorCode":    "401",
				"errorMessage": "请先登录！",
				"success":      true,
			})
		} else {
			c.JSON(http.StatusAccepted, gin.H{
				"name":   "Sa1ka",
				"userid": "00000001",
			})
		}
	})

	router.POST("/api/login/account", func(c *gin.Context) {
		session := sessions.Default(c)

		var loginInfo user
		if err := c.BindJSON(&loginInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "parse heartbeat from http post request error",
			})
			return
		}

		if loginInfo.Username == "admin" && loginInfo.Password == "admin" {
			session.Set("User", "admin")
			session.Save()

			c.JSON(http.StatusOK, gin.H{
				"status":           "ok",
				"type":             loginInfo.Type,
				"currentAuthority": "admin",
			})

			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":           "error",
			"type":             loginInfo.Type,
			"currentAuthority": "guest",
		})

	})

	router.POST("/api/login/outLogin", func(c *gin.Context) {
		session := sessions.Default(c)

		session.Delete("User")
		session.Save()

		c.JSON(http.StatusOK, gin.H{
			"data":    gin.H{},
			"success": true,
		})

	})

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
