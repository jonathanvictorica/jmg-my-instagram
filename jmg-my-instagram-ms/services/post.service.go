package services

type IPostService interface {
	CreatePost(postRequest *PostRequest) (uint, error)
	GetPostsByUser(userID string) ([]PostResponse, error)
}

type PostRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	Sections    []Section `json:"sections"`
}

type Section struct {
	Images      uint     `json:"multimedia_content_id"`
	Description string   `json:"description"`
	Hashtags    []string `json:"hashtags"`
	Users       []uint   `json:"users"`
	Location    uint     `json:"location"`
}

type PostResponse struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
}
