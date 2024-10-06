package service

import (
	"crypto/sha1"
	"fmt"
	
	todo "github.com/kingxl111/RESTapiService"
	"github.com/kingxl111/RESTapiService/pkg/repository"
)

const salt = "erkqwekldfsl12378ds8qke32" // соль для хэширования пароля

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// Еще на слой ниже 

func (s* AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}