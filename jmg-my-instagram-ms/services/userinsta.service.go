package services

type IUserService interface {
	SearchUsers(name string, top int) ([]UserResponse, error)
	GetUserProfile(userID string) (*UserProfileResponse, error)
}

type UserResponse struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}

type UserProfileResponse struct {
	Username    string `json:"username"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}
