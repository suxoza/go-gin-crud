package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/suxoza/go-gin-crud/initializers"
	"github.com/suxoza/go-gin-crud/models"
)

var reqeustBody struct {
	Title string
	Body  string
}

func PostsCreate(c *gin.Context) {

	c.Bind(&reqeustBody)

	post := models.Post{Title: reqeustBody.Title, Body: reqeustBody.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		log.Fatal("Can`t create post")
		return
	}
	c.JSON(200, gin.H{
		"data": post,
	})
}

func SetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}

	c.Bind(&reqeustBody)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: reqeustBody.Title,
		Body:  reqeustBody.Body,
	})

	c.JSON(200, gin.H{
		"data": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)
	c.Status(200)
}

func PostsList(c *gin.Context) {

	var posts []models.Post
	if err := initializers.DB.Find(&posts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{"data": posts})
	}
}

func GetPostById(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{"data": post})
	}
}
