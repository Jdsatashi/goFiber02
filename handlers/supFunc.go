package handlers

import (
	"github.com/Jdsatashi/goFiber02/models"
	repo "github.com/Jdsatashi/goFiber02/repositories"
)

type UserHandler struct {
}
type BookHandler struct{}

func (u *UserHandler) ToUserResponse(user models.Users) *repo.UserResponse {
	return &repo.UserResponse{
		UserCode: user.UserCode,
		Name:     user.Username,
		Email:    user.Email,
	}
}

func (u *UserHandler) ToUsersResponse(users []models.Users) []*repo.UserResponse {
	var responses []*repo.UserResponse
	for _, user := range users {
		responses = append(responses, u.ToUserResponse(user))
	}
	return responses
}

func (b *BookHandler) ToBookResponse(book models.Books) *repo.BookResponse {
	return &repo.BookResponse{
		ID:        book.ID,
		Author:    book.Author,
		Title:     book.Title,
		Publisher: book.Publisher,
		WriterID:  book.WriterID,
	}
}

func (b *BookHandler) ToBooksResponse(books []models.Books) []*repo.BookResponse {
	var responses []*repo.BookResponse
	for _, book := range books {
		responses = append(responses, b.ToBookResponse(book))
	}
	return responses
}
