package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josh1248/forum-website-backend/internal/entities"
	"github.com/josh1248/forum-website-backend/internal/models"
)

func GetAllPosts(c *gin.Context) {
	posts, err := models.FindAllPosts()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, posts)
	}
}

// weigh server-side vs client-side data validation. security vs ease of implementation.
func CreatePost(c *gin.Context) {
	var newPost entities.Post

	//Unmarshals JSON data and reads into struct
	err := c.ShouldBindJSON(&newPost)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}

	err = models.CreatePost(newPost)
	if err != nil {
		//placeholder controls.
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, "successfully created.")
		return
	}
}

func DeletePost(c *gin.Context) {

}
