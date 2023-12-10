package main

import (
	"github.com/suxoza/go-gin-crud/initializers"
	"github.com/suxoza/go-gin-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
