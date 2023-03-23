package controllers

import (
	"github.com/Vrivaans/go-crud/initializers"
	"github.com/Vrivaans/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Crea un Posteo
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//Retorna

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	//Obtener los posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Responder los posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}
