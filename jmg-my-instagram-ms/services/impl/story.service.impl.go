package impl

import (
	"jmg-my-instagram-ms/repository"
	"jmg-my-instagram-ms/services"
)

type StoryService struct {
	storyRepository       repository.IStoryRepository
	locationRepo          repository.ILocationRepository
	multimediaContentRepo repository.IMultimediaContentRepository
	hashtagRepo           repository.IHashtagRepository
}

func NewStoryService(storyRepository repository.IStoryRepository,
	locationRepo repository.ILocationRepository,
	multimediaContentRepo repository.IMultimediaContentRepository,
	hashtagRepo repository.IHashtagRepository) *StoryService {
	return &StoryService{
		storyRepository:       storyRepository,
		locationRepo:          locationRepo,
		multimediaContentRepo: multimediaContentRepo,
		hashtagRepo:           hashtagRepo,
	}
}

func (s *StoryService) CreateStory(historyRequest *services.StoryRequest) (string, error) {
	return "", nil
}
func (s *StoryService) GetStoriesByUser(userID string) ([]services.StoryResponse, error) {
	return nil, nil
}
func (s *StoryService) AddContentToStory(historyID string, contentRequest *services.StoryContentRequest) (string, error) {
	return "", nil
}
