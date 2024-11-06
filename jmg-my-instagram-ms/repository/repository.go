package repository

import (
	models "jmg-my-instagram-ms/model"
)

// Interfaz para Hashtag
type IHashtagRepository interface {
	Save(hashtag *models.Hashtag) error
	Update(hashtag *models.Hashtag) error
	Delete(id uint) error
	FindById(id uint) (*models.Hashtag, error)
	FindByName(name string) (*models.Hashtag, error)
	FindAll() ([]*models.Hashtag, error)
}

// Interfaz para Location
type ILocationRepository interface {
	Save(location *models.Location) error
	Update(location *models.Location) error
	Delete(id uint) error
	FindById(id uint) (*models.Location, error)
	FindAll() ([]*models.Location, error)
	Search(name string, top int) ([]models.Location, error)
}

// Interfaz para MultimediaContent
type IMultimediaContentRepository interface {
	Save(multimediaContent *models.MultimediaContent) error
	Update(multimediaContent *models.MultimediaContent) error
	Delete(id uint) error
	FindById(id uint) (*models.MultimediaContent, error)
	FindAll() ([]*models.MultimediaContent, error)
}

// Interfaz para UserInsta
type IUserInstaRepository interface {
	Save(user *models.UserInsta) error
	Update(user *models.UserInsta) error
	Delete(id uint) error
	FindById(id uint) (*models.UserInsta, error)
	FindByName(username string) (*models.UserInsta, error)
	FindAll() ([]*models.UserInsta, error)
	Search(name string, top int) ([]models.UserInsta, error)
}

// Interfaz para Post
type IPostRepository interface {
	Save(post *models.Post) error
	Update(post *models.Post) error
	Delete(id uint) error
	FindById(id uint) (*models.Post, error)
	FindAll() ([]*models.Post, error)
}

// Interfaz para PostContent
type IPostContentRepository interface {
	Save(postContent *models.PostContent) error
	Update(postContent *models.PostContent) error
	Delete(id uint) error
	FindById(id uint) (*models.PostContent, error)
	FindAll() ([]*models.PostContent, error)
}

// Interfaz para PostContentHashtag
type IPostContentHashtagRepository interface {
	Save(postContentHashtag *models.PostContentHashtag) error
	Update(postContentHashtag *models.PostContentHashtag) error
	Delete(id uint) error
	FindById(id uint) (*models.PostContentHashtag, error)
	FindAll() ([]*models.PostContentHashtag, error)
}

// Interfaz para PostContentUsers
type IPostContentUsersRepository interface {
	Save(postContentUser *models.PostContentUsers) error
	Update(postContentUser *models.PostContentUsers) error
	Delete(id uint) error
	FindById(id uint) (*models.PostContentUsers, error)
	FindAll() ([]*models.PostContentUsers, error)
}

// Interfaz para Story
type IStoryRepository interface {
	Save(history *models.Story) error
	Update(history *models.Story) error
	Delete(id uint) error
	FindById(id uint) (*models.Story, error)
	FindAll() ([]*models.Story, error)
}

// Interfaz para StoryContent
type IStoryContentRepository interface {
	Save(historyContent *models.StoryContent) error
	Update(historyContent *models.StoryContent) error
	Delete(id uint) error
	FindById(id uint) (*models.StoryContent, error)
	FindAll() ([]*models.StoryContent, error)
}

// Interfaz para StoryContentHashtag
type IStoryContentHashtagRepository interface {
	Save(historyContentHashtag *models.StoryContentHashtag) error
	Update(historyContentHashtag *models.StoryContentHashtag) error
	Delete(id uint) error
	FindById(id uint) (*models.StoryContentHashtag, error)
	FindAll() ([]*models.StoryContentHashtag, error)
}

// Interfaz para StoryContentUsers
type IStoryContentUsersRepository interface {
	Save(historyContentUser *models.StoryContentUsers) error
	Update(historyContentUser *models.StoryContentUsers) error
	Delete(id uint) error
	FindById(id uint) (*models.StoryContentUsers, error)
	FindAll() ([]*models.StoryContentUsers, error)
}
