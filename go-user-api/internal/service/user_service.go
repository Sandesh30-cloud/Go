package service

import "github.com/Sandesh30-cloud/go-user-api/internal/model"

type UserService struct {
	users []model.User
}

func NewUserService() *UserService {
	return &UserService{
		users: []model.User{},
	}
}

func (s *UserService) GetAll() []model.User {
	return s.users
}

func (s *UserService) GetByID(id int) *model.User {
	for _, user := range s.users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func (s *UserService) Create(user model.User) {
	s.users = append(s.users, user)
}

func (s *UserService) Delete(id int) bool {
	for i, user := range s.users {
		if user.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return true
		}
	}
	return false
}
