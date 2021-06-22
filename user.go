package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	userModel struct {
		gorm.Model
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	transformedUser struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}
)

func createUser(c *gin.Context) {
	user := userModel{Name: c.PostForm("name"), Password: c.PostForm("password")}
	db.Save(&user)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User item created successfully!", "resourceId": user.ID})
}

func fetchAllUsers(c *gin.Context) {
	var users []userModel
	var _users []transformedUser
	db.Find(&users)
	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	for _, item := range users {
		_users = append(_users, transformedUser{ID: item.ID, Name: item.Name, Password: item.Password})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _users})
}

func fetchUser(c *gin.Context) {
	var user userModel
	userId := c.Param("id")
	db.First(&user, userId)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}
	_user := transformedUser{ID: user.ID, Name: user.Name}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _user})
}

func updateUser(c *gin.Context) {
	var user userModel
	userId := c.Param("id")
	db.First(&user, userId)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	db.Model(&user).Update("name", c.PostForm("name")).Update("password", c.PostForm("password"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User updated successfully!"})
}

func deleteUser(c *gin.Context) {
	var user userModel
	userId := c.Param("id")
	db.First(&user, userId)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User deleted successfully!"})
}
