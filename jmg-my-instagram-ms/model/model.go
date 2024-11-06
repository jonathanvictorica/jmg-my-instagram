package models

import (
	"time"
)

// Hashtag
type Hashtag struct {
	HashtagID uint   `gorm:"primaryKey;column:hashtag_id;not null" json:"hashtag_id"`
	Name      string `gorm:"column:name;not null;unique" json:"name"`
}

func (Hashtag) TableName() string {
	return "hashtags"
}

// Location
type Location struct {
	LocationID     uint   `gorm:"primaryKey;column:location_id;not null" json:"location_id"`
	Country        string `gorm:"column:country;not null" json:"country"`
	City           string `gorm:"column:city;not null" json:"city"`
	InfoAdditional string `gorm:"column:info_aditional" json:"info_aditional,omitempty"`
}

func (Location) TableName() string {
	return "locations"
}

// MultimediaContent
type MultimediaContent struct {
	MultimediaContentID uint   `gorm:"primaryKey;column:multimedia_content_id;not null" json:"multimedia_content_id"`
	ContentURL          string `gorm:"column:content_url;not null" json:"content_url"`
	ContentType         string `gorm:"column:content_type;not null;check:content_type IN ('IMAGE', 'VIDEO')" json:"content_type"`
}

func (MultimediaContent) TableName() string {
	return "multimedia_contents"
}

// UserInsta
type UserInsta struct {
	UserID                     uint   `gorm:"primaryKey;column:user_id;not null" json:"user_id"`
	Username                   string `gorm:"column:username;not null;unique" json:"username"`
	Email                      string `gorm:"column:email;not null;unique" json:"email"`
	ProfileMultimediaContentID *uint  `gorm:"column:profile_multimedia_content_id" json:"profile_multimedia_content_id,omitempty"`
	Description                string `gorm:"column:description" json:"description,omitempty"`
}

func (UserInsta) TableName() string {
	return "user_instas"
}

// Post
type Post struct {
	PostID    uint      `gorm:"primaryKey;column:post_id;not null" json:"post_id"`
	UserID    uint      `gorm:"column:user_id;not null" json:"user_id"`
	Title     string    `gorm:"column:title" json:"title"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	PostType  string    `gorm:"column:post_type;not null;check:post_type IN ('POST', 'REEL')" json:"post_type"`
}

func (Post) TableName() string {
	return "post"
}

// PostContent
type PostContent struct {
	PostContentID uint `gorm:"primaryKey;column:post_content_id;not null" json:"post_content_id"`

	Description string `gorm:"column:description" json:"description,omitempty"`
	Order       int    `gorm:"column:orden" json:"orden,omitempty"`
	Enabled     bool   `gorm:"column:enabled;default:true" json:"enabled"`

	// Foreign key fields
	PostID              uint `gorm:"column:post_id;not null" json:"post_id"`                    // Clave foránea para Post
	MultimediaContentID uint `gorm:"column:multimedia_content_id" json:"multimedia_content_id"` // Clave foránea para MultimediaContent
	LocationID          uint `gorm:"column:location_id" json:"location_id"`                     // Clave foránea para Location

	Post              *Post              `gorm:"foreignKey:PostID" json:"post,omitempty"`
	MultimediaContent *MultimediaContent `gorm:"foreignKey:MultimediaContentID" json:"multimedia_content,omitempty"`
	Location          *Location          `gorm:"foreignKey:LocationID" json:"location,omitempty"`
}

func (PostContent) TableName() string {
	return "post_content"
}

// PostContentHashtag
type PostContentHashtag struct {
	PostContentHashtagID uint `gorm:"primaryKey;column:post_content_hashtag_id;not null" json:"post_content_hashtag_id"`
	PostContentID        uint `gorm:"column:post_content_id;not null" json:"post_content_id"`
	HashtagID            uint `gorm:"column:hashtag_id;not null" json:"hashtag_id"`
}

func (PostContentHashtag) TableName() string {
	return "post_content_hashtags"
}

// PostContentUsers
type PostContentUsers struct {
	PostContentUsersID uint `gorm:"primaryKey;column:post_content_users_id;not null" json:"post_content_users_id"`
	PostContentID      uint `gorm:"column:post_content_id;not null" json:"post_content_id"`
	UserID             uint `gorm:"column:user_id;not null" json:"user_id"`
}

func (PostContentUsers) TableName() string {
	return "post_content_users"
}

// Story
type Story struct {
	StoryID             uint               `gorm:"primaryKey;column:story_id;not null" json:"story_id"`
	Title               string             `gorm:"column:title" json:"title"`
	MultimediaContentID uint               `gorm:"column:multimedia_content_id;not null" json:"multimedia_content_id"`
	Order               int                `gorm:"column:orden;not null" json:"orden"`
	MultimediaContent   *MultimediaContent `gorm:"foreignKey:MultimediaContentID" json:"multimedia_content,omitempty"`
}

// StoryContent
type StoryContent struct {
	StoryContentID      uint               `gorm:"primaryKey;column:story_content_id;not null" json:"story_content_id"`
	StoryID             uint               `gorm:"column:story_id;not null" json:"story_id"`
	MultimediaContentID uint               `gorm:"column:multimedia_content_id;not null" json:"multimedia_content_id"`
	LocationID          *uint              `gorm:"column:location_id" json:"location_id,omitempty"`
	Description         string             `gorm:"column:description" json:"description,omitempty"`
	Order               int                `gorm:"column:orden" json:"orden,omitempty"`
	Enabled             bool               `gorm:"column:enabled;default:true" json:"enabled"`
	MultimediaContent   *MultimediaContent `gorm:"foreignKey:MultimediaContentID" json:"multimedia_content,omitempty"`
	Location            *Location          `gorm:"foreignKey:LocationID" json:"location,omitempty"`
}

// StoryContentHashtag
type StoryContentHashtag struct {
	StoryContentHashtagID uint `gorm:"primaryKey;column:story_content_hashtag_id;not null" json:"story_content_hashtag_id"`
	StoryContentID        uint `gorm:"column:story_content_id;not null" json:"story_content_id"`
	HashtagID             uint `gorm:"column:hashtag_id;not null" json:"hashtag_id"`
}

// StoryContentUsers
type StoryContentUsers struct {
	StoryContentUsersID uint `gorm:"primaryKey;column:story_content_users_id;not null" json:"story_content_users_id"`
	StoryContentID      uint `gorm:"column:story_content_id;not null" json:"story_content_id"`
	UserID              uint `gorm:"column:user_id;not null" json:"user_id"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
