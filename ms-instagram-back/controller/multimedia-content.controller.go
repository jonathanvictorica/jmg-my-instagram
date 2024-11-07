package controller

import (
	"jmg-my-instagram-ms/services"
)

type MultimediaContentController struct {
	multimediaContentService services.IMultimediaContentService
}

func NewMultimediaContentController(
	multimediaContentService services.IMultimediaContentService) MultimediaContentController {
	return MultimediaContentController{
		multimediaContentService: multimediaContentService,
	}
}
