package service

import (
	"crypto/sha1"
	"fmt"
	"time"
	"todo"
	"todo/pkg/repository"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "sadasdasdsd"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

//Создаем метод объекта который получает данные в виде todo.User и дальше уже сам работает с репозиторием
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	//проверяем пользователя
	id, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	//если все ок то генерируем токе
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			//время действия токена
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			//время создания токена
			IssuedAt: time.Now().Unix(),
		},
		//то что мы шифруем в токене
		id,
	})

	//возращаем токен
	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
