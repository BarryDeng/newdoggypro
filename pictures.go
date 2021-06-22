package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	pictureModel struct {
		gorm.Model
		Content string `json:"content"`
	}

	transformedPicture struct {
		ID      uint   `json:"id"`
		Content string `json:"content"`
	}
)

func createPicture(c *gin.Context) {
	pic := pictureModel{Content: c.PostForm("content")}
	db.Save(&pic)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Picture item created successfully!", "resourceId": pic.ID})
}

func fetchAllPictures(c *gin.Context) {
	var pics []pictureModel
	var _pics []transformedPicture
	db.Find(&pics)
	if len(pics) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	for _, item := range pics {
		_pics = append(_pics, transformedPicture{ID: item.ID, Content: item.Content})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _pics})
}

func fetchPicture(c *gin.Context) {
	var pic pictureModel
	picId := c.Param("id")
	db.First(&pic, picId)
	if pic.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No pic found!"})
		return
	}
	_pic := transformedPicture{ID: pic.ID, Content: pic.Content}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _pic})
}

func updatePicture(c *gin.Context) {
	var pic pictureModel
	picId := c.Param("id")
	db.First(&pic, picId)
	if pic.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	db.Model(&pic).Update("content", c.PostForm("content"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Pic updated successfully!"})
}

func deletePicture(c *gin.Context) {
	var pic pictureModel
	picId := c.Param("id")
	db.First(&pic, picId)
	if pic.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	db.Delete(&pic)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Pic deleted successfully!"})
}
