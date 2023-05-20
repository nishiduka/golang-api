package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nishiduka/rest-api/initializers"
	"github.com/nishiduka/rest-api/models"
)

func PostDetails(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	if post.ID == 0 {
		c.JSON(400, gin.H{
			"data": "",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostsList(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"data": posts,
	})
}

func PostCreate(c *gin.Context) {
	var newPost models.Post

	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Fields",
		})
		return
	}

	result := initializers.DB.Create(&newPost)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Error on create",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var request models.Post
	c.BindJSON(&request)

	var post models.Post

	initializers.DB.First(&post, id)

	request.ID = post.ID
	initializers.DB.Updates(&request)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"message": "success",
	})
}
