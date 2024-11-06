package impl

import (
	"jmg-my-instagram-ms/repository"
	"jmg-my-instagram-ms/services"
)

type MultimediaContentService struct {
	multimediaContentRepo repository.IMultimediaContentRepository
}

func NewMultimediaContentService(multimediaContentRepo repository.IMultimediaContentRepository) *MultimediaContentService {
	return &MultimediaContentService{multimediaContentRepo}
}

func (s *MultimediaContentService) CreateMultimediaContent(multimediaContentRequest *services.MultimediaContentRequest) (string, error) {
	return "", nil
}
