package factory

import (
	"jmg-my-instagram-ms/controller"
	"jmg-my-instagram-ms/repository"
	"jmg-my-instagram-ms/services"
	"jmg-my-instagram-ms/services/impl"
)

type Container struct {
	HashtagRepo             repository.IHashtagRepository
	LocationRepo            repository.ILocationRepository
	MultimediaContentRepo   repository.IMultimediaContentRepository
	UserInstaRepo           repository.IUserInstaRepository
	PostRepo                repository.IPostRepository
	PostContentRepo         repository.IPostContentRepository
	PostContentHashtagRepo  repository.IPostContentHashtagRepository
	PostContentUsersRepo    repository.IPostContentUsersRepository
	StoryRepo               repository.IStoryRepository
	StoryContentRepo        repository.IStoryContentRepository
	StoryContentHashtagRepo repository.IStoryContentHashtagRepository
	StoryContentUsersRepo   repository.IStoryContentUsersRepository

	LocationService          services.ILocationService
	MultimediaContentService services.IMultimediaContentService
	PostService              services.IPostService
	UserStoryService         services.IUserStoryService
	UserService              services.IUserService

	LocationController          controller.LocationController
	MultimediaContentController controller.MultimediaContentController
	PostController              controller.PostController
	StoryController             controller.StoryController
	UserInstagramController     controller.UserInstagramController
}

var globalContainer *Container

func FactoryContainer() *Container {
	if globalContainer == nil {
		globalContainer = newContainer()
	}
	return globalContainer
}

func newContainer() *Container {

	globalContainer = &Container{}
	// Repository
	globalContainer.HashtagRepo = repository.NewHashtagRepository()
	globalContainer.LocationRepo = repository.NewLocationRepository()
	globalContainer.MultimediaContentRepo = repository.NewMultimediaContentRepository()
	globalContainer.UserInstaRepo = repository.NewUserInstaRepository()
	globalContainer.PostRepo = repository.NewPostRepository()
	globalContainer.PostContentRepo = repository.NewPostContentRepository()
	globalContainer.PostContentHashtagRepo = repository.NewPostContentHashtagRepository()
	globalContainer.PostContentUsersRepo = repository.NewPostContentUsersRepository()
	globalContainer.StoryRepo = repository.NewStoryRepository()
	globalContainer.StoryContentRepo = repository.NewStoryContentRepository()
	globalContainer.StoryContentHashtagRepo = repository.NewStoryContentHashtagRepository()
	globalContainer.StoryContentUsersRepo = repository.NewStoryContentUsersRepository()

	// Services
	globalContainer.LocationService = impl.NewLocationService(globalContainer.LocationRepo)
	globalContainer.MultimediaContentService = impl.NewMultimediaContentService(globalContainer.MultimediaContentRepo)
	globalContainer.UserService = impl.NewUserInstaService(globalContainer.UserInstaRepo)
	globalContainer.PostService = impl.NewPostService(
		globalContainer.PostContentHashtagRepo,
		globalContainer.PostContentUsersRepo,
		globalContainer.PostRepo,
		globalContainer.LocationRepo,
		globalContainer.MultimediaContentRepo,
		globalContainer.HashtagRepo,
		globalContainer.UserInstaRepo, globalContainer.PostContentRepo)
	globalContainer.UserStoryService = impl.NewStoryService(globalContainer.StoryRepo, globalContainer.LocationRepo, globalContainer.MultimediaContentRepo, globalContainer.HashtagRepo)
	globalContainer.LocationController = controller.NewLocationController(globalContainer.LocationService)
	globalContainer.MultimediaContentController = controller.NewMultimediaContentController(globalContainer.MultimediaContentService)
	globalContainer.PostController = controller.NewPostController(globalContainer.PostService)
	globalContainer.StoryController = controller.NewStoryController(globalContainer.UserStoryService)
	globalContainer.UserInstagramController = controller.NewUserInstagramController(globalContainer.UserService)

	return globalContainer
}
