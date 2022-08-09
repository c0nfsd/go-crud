package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data from the request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id from url
	id:= c.Param("id")

	//Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})

}
