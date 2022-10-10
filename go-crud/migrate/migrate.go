package main

import (
	"github.com/hambalee/go-crud/initializers"
	"github.com/hambalee/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
