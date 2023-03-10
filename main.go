package main

import (
	"github.com/Vrivaans/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola mundo",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
