package handlers

import (
	"github.com/Jdsatashi/goFiber02/models"
	repo "github.com/Jdsatashi/goFiber02/repositories"
)

type UserHandler struct {
}

func (u *UserHandler) ToUserResponse(user models.Users) *repo.UserResponse {
	return &repo.UserResponse{
		UserCode: user.UserCode,
		Name:     user.Username,
		Email:    user.Email,
	}
}

//func (u *UserHandler) ToUsersResponse(users []*models.Users) []*repo.UserResponse {
//	var responseList []*repo.UserResponse
//	for _, user := range users {
//		responseList = append(responseList, u.ToUserResponse(user))
//	}
//	return responseList
//}

func (u *UserHandler) ToUsersResponse(users []models.Users) []*repo.UserResponse {
	var responses []*repo.UserResponse
	for _, user := range users {
		responses = append(responses, u.ToUserResponse(user))
	}
	return responses
}

//response := &repo.UserResponse{
//	ID:    user.ID,
//	Name:  user.Username,
//	Email: user.Email,
//}
//responses = append(responses, response)
