package impl

import (
	"jmg-my-instagram-ms/repository"
	"jmg-my-instagram-ms/services"
)

type UserInstaService struct {
	userInstaRepo repository.IUserInstaRepository
}

func NewUserInstaService(userInstaRepo repository.IUserInstaRepository) *UserInstaService {
	return &UserInstaService{userInstaRepo: userInstaRepo}
}

func (s *UserInstaService) SearchUsers(name string, top int) ([]services.UserResponse, error) {
	users, _ := s.userInstaRepo.Search(name, top)

	var result = make([]services.UserResponse, 0)
	for _, user := range users {
		result = append(result, services.UserResponse{
			UserID:   user.UserID,
			Username: user.Username,
		})
	}
	return result, nil
}
func (s *UserInstaService) GetUserProfile(userID string) (*services.UserProfileResponse, error) {
	return nil, nil
}
