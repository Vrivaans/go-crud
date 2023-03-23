package controllers

import (
	"github.com/Vrivaans/go-crud/initializers"
	"github.com/Vrivaans/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body

	//Crea un Posteo
	post := models.Post{Title: "Hola mundo", Body: "Hola mundo desde Go, Gin, GORM"}
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
