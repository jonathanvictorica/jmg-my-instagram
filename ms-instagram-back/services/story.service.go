package services

type IUserStoryService interface {
	CreateStory(historyRequest *StoryRequest) (string, error)
	GetStoriesByUser(userID string) ([]StoryResponse, error)
	AddContentToStory(historyID string, contentRequest *StoryContentRequest) (string, error)
}

type StoryResponse struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
}

type StoryRequest struct {
	UserID   string `json:"user_id"`
	ImageURL string `json:"image_url"`
	Title    string `json:"title"`
}

type StoryContentRequest struct {
	UserID     string   `json:"user_id"`
	ImageURL   string   `json:"image_url"`
	Users      []string `json:"users"`
	LocationID string   `json:"location_id,omitempty"`
	Hashtags   []string `json:"hashtags"`
}
