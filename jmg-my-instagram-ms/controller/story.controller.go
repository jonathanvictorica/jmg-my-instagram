package controller

import (
	"github.com/gin-gonic/gin"
	"jmg-my-instagram-ms/services"
)

type StoryController struct {
	storyService services.IUserStoryService
}

func (c StoryController) GetByUser(context *gin.Context) {

}

func (c StoryController) Create(context *gin.Context) {

}

func NewStoryController(storyService services.IUserStoryService) StoryController {
	return StoryController{storyService}
}
