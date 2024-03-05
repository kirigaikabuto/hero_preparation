package service

import (
	"github.com/kirigaikabuto/hero_preparation/projects/project1/models"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/pkg/repository"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/utils"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = utils.GetHashedPassword(user.Password)
	return a.repo.CreateUser(user)
}
