package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hackjakarta2024/backend/internal/model"
	"github.com/hackjakarta2024/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type userService struct {
	userRepo repository.UserRepository
}

type UserService interface {
	CreateUser(user model.User) error
	Login(email, password string) (string, error)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(user model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.ID = uuid.New()

	if err := s.userRepo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRETKEY")))
	if err != nil {
		log.Println("error: " + err.Error())
		return "", err
	}

	return tokenString, nil
}
