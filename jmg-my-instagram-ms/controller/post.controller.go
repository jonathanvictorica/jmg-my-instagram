package controller

import (
	"github.com/gin-gonic/gin"
	"jmg-my-instagram-ms/services"
	"net/http"
)

type PostController struct {
	postService services.IPostService
}

func (c PostController) Create(context *gin.Context) {
	var request services.PostRequest
	if err := context.BindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postId, err := c.postService.CreatePost(&request)
	if err != nil {
		return
	}
	context.JSON(http.StatusCreated, gin.H{"post_id": postId})
}

func (c PostController) GetByUser(context *gin.Context) {

}

func NewPostController(postService services.IPostService) PostController {
	return PostController{postService: postService}
}
