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
	r.GET("/post/:id", controllers.PostsShow)
	r.PUT("post/update/:id", controllers.PostUpdate)
	r.DELETE("/delete/:id", controllers.PostDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
