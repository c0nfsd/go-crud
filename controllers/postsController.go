package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

/*
	Get data from the request body
	create a post
	return it
*/

func PostsCreate(c *gin.Context) {

	post := models.Post{Title: "First", Body: "Post body"}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}
