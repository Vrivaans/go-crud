package main

import (
	"github.com/Vrivaans/go-crud/controllers"
	"github.com/Vrivaans/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/post", controllers.PostsCreate)
	r.GET("/", controllers.PostIndex)
	r.Run() // listen and serve on 0.0.0.0:8080
}
