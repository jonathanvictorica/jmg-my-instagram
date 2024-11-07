package impl

import (
	"jmg-my-instagram-ms/exception"
	models "jmg-my-instagram-ms/model"
	"jmg-my-instagram-ms/repository"
	"jmg-my-instagram-ms/services"
	"time"
)

type PostService struct {
	postContentHashtagRepository repository.IPostContentHashtagRepository
	postContentUsersRepository   repository.IPostContentUsersRepository
	postContentRepository        repository.IPostContentRepository
	postRepository               repository.IPostRepository
	locationRepo                 repository.ILocationRepository
	multimediaContentRepo        repository.IMultimediaContentRepository
	hashtagRepo                  repository.IHashtagRepository
	userRepo                     repository.IUserInstaRepository
}

func NewPostService(
	postContentHashtagRepository repository.IPostContentHashtagRepository,
	postContentUsersRepository repository.IPostContentUsersRepository,
	postRepository repository.IPostRepository,
	locationRepo repository.ILocationRepository,
	multimediaContentRepo repository.IMultimediaContentRepository,
	hashtagRepo repository.IHashtagRepository,
	userRepo repository.IUserInstaRepository,
	postContentRepository repository.IPostContentRepository) *PostService {
	return &PostService{
		postContentHashtagRepository: postContentHashtagRepository,
		postContentUsersRepository:   postContentUsersRepository,
		postRepository:               postRepository,
		locationRepo:                 locationRepo,
		multimediaContentRepo:        multimediaContentRepo,
		hashtagRepo:                  hashtagRepo,
		userRepo:                     userRepo,
		postContentRepository:        postContentRepository,
	}
}

func (s *PostService) CreatePost(postRequest *services.PostRequest) (uint, error) {
	var post = new(models.Post)
	post.CreatedAt = time.Now()
	post.PostType = "POST"
	post.Title = postRequest.Title
	post.UserID = postRequest.UserID

	user, _ := s.userRepo.FindById(postRequest.UserID)
	if user == nil {
		return 0, exception.UserNotFount
	}

	err := s.postRepository.Save(post)
	if err != nil {
		return 0, err
	}

	// Sections
	for orden, unit := range postRequest.Sections {
		var postContent = new(models.PostContent)
		postContent.PostID = post.PostID

		location, err := s.locationRepo.FindById(unit.Location)
		if location == nil {
			return 0, exception.LocationNotFound
		}
		postContent.LocationID = location.LocationID

		multimedia, err := s.multimediaContentRepo.FindById(unit.Images)
		if multimedia == nil {
			return 0, exception.MultimediaNotFound
		}
		postContent.MultimediaContentID = multimedia.MultimediaContentID

		postContent.Enabled = true
		postContent.Description = unit.Description
		postContent.Order = orden

		err = s.postContentRepository.Save(postContent)
		if err != nil {
			return 0, err
		}

		for _, hashtagName := range unit.Hashtags {

			hashtag, _ := s.hashtagRepo.FindByName(hashtagName)
			if hashtag == nil {
				m := models.Hashtag{Name: hashtagName}
				s.hashtagRepo.Save(&m)
				hashtag, _ = s.hashtagRepo.FindByName(hashtagName)
			}

			var postHashtag = new(models.PostContentHashtag)
			postHashtag.PostContentID = postContent.PostContentID
			postHashtag.HashtagID = hashtag.HashtagID
			err := s.postContentHashtagRepository.Save(postHashtag)
			if err != nil {
				return 0, err
			}

		}

		for _, userEtiqueted := range unit.Users {
			userEtiquetedPost, _ := s.userRepo.FindById(userEtiqueted)
			if userEtiquetedPost == nil {
				return 0, exception.UserNotFount
			}

			userPost := models.PostContentUsers{
				PostContentUsersID: 0,
				PostContentID:      postContent.PostContentID,
				UserID:             userEtiquetedPost.UserID,
			}

			err := s.postContentUsersRepository.Save(&userPost)
			if err != nil {
				return 0, err
			}
		}
	}

	return post.PostID, nil
}
func (s *PostService) GetPostsByUser(userID string) ([]services.PostResponse, error) {
	return nil, nil
}
