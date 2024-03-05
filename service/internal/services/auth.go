package services

import (
	"TrafficPolice/internal/domain"
	"TrafficPolice/internal/repository"
	"TrafficPolice/internal/tokens"
	"TrafficPolice/pkg/hash"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type AuthService interface {
	RegisterExpert(input domain.User) error
	SignIn(input domain.User) (string, error)
	ConfirmExpert(data domain.ConfirmExpert) error
	ParseAccessToken(accessToken string) (tokens.TokenInfo, error)
}

type authService struct {
	repo           repository.AuthRepo
	hasher         hash.PasswordHasher
	tokenManager   tokens.TokenManager
	accessTokenTTL time.Duration
}

func NewAuthService(repo repository.AuthRepo, tokenManager tokens.TokenManager) AuthService {
	return &authService{
		repo:           repo,
		hasher:         hash.NewSHA1Hasher("salt"),
		tokenManager:   tokenManager,
		accessTokenTTL: 30 * 24 * time.Hour,
	}
}

func (s *authService) RegisterExpert(user domain.User) error {
	err := s.repo.CheckUserExists(user.Username)
	if err == nil {
		return fmt.Errorf("user with username '%s' already exists", user.Username)
	}

	hashedPass, err := s.hasher.Hash(user.Password)
	if err != nil {
		return err
	}

	user.ID = uuid.New()
	user.Password = hashedPass
	user.UserRole = "expert"
	err = s.repo.InsertUser(user)
	if err != nil {
		return err
	}

	err = s.repo.InsertExpert(domain.Expert{
		ID:   uuid.New(),
		User: user,
	})

	return err
}

func (s *authService) SignIn(input domain.User) (string, error) {
	user, err := s.repo.SignIn(input.Username)
	if err != nil {
		return "", err
	}

	inputHashPass, err := s.hasher.Hash(input.Password)
	if err != nil {
		return "", err
	}

	if user.Password != inputHashPass {
		return "", fmt.Errorf("invalid password")
	}

	return s.tokenManager.NewJWT(tokens.TokenInfo{
		UserID:   user.ID.String(),
		UserRole: domain.Role(user.UserRole),
	}, s.accessTokenTTL)
}

func (s *authService) ConfirmExpert(data domain.ConfirmExpert) error {
	return s.repo.ConfirmExpert(data)
}

func (s *authService) ParseAccessToken(accessToken string) (tokens.TokenInfo, error) {
	return s.tokenManager.Parse(accessToken)
}