package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suxoza/go-gin-crud/controllers"
	"github.com/suxoza/go-gin-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.GET("/posts", controllers.PostsList)
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts/:id", controllers.GetPostById)
	r.PUT("/posts/:id", controllers.SetPost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run()
}
