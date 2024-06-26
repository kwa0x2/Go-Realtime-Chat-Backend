package service

import (
	"github.com/kwa0x2/realtime-chat-backend/models"
	"github.com/kwa0x2/realtime-chat-backend/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (s *UserService) IsUsernameUnique(username string) bool {
	return s.UserRepository.IsUsernameUnique(username)
}





func (s *UserService) InsertUser(user *models.User) error{
	return s.UserRepository.InsertUser(user)
}