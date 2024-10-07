package service

import (
	"crypto/sha1"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	todo "github.com/kingxl111/RESTapiService"
	"github.com/kingxl111/RESTapiService/pkg/repository"
)

const (
	salt = "erkqwekldfsl12378ds8qke32" // соль для хэширования пароля
	signingKey = "skdf12njx83xu39ck2d91cwjel" // для подписания jwt токенов 	
)

type tokenClaims struct {
	jwt.RegisteredClaims //Встраиваем тип. Здесь стандарты Claims
	UserId int `json:"userId"`
}

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

func (s* AuthService) GenerateToken(username, password string) (string, error) {
	// get user from db
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		UserId: user.Id,
	})
	
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(signingKey), nil
	})
	

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok || claims == nil {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

