package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/olegvolkov91/go-todo-rest-tutorial"
	"github.com/olegvolkov91/go-todo-rest-tutorial/package/repository"
)

const salt string = "1klnld90131jDS#SL!S(($"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
