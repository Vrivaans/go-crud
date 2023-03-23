package main

import (
	"github.com/Vrivaans/go-crud/initializers"
	"github.com/Vrivaans/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
