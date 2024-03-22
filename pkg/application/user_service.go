package application

import "github.com/tutorial_ddd_go-mongodb/pkg/domain"

type UserService struct {
	Repository domain.UserRepository
}

func (s *UserService) GetUser(id string) (domain.User, error) {
	return s.Repository.GetUser(id)
}

func (s *UserService) CreateUser(user domain.User) (string, error) {
	return s.Repository.CreateUser(user)
}
