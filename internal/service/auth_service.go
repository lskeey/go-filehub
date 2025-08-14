package service

import (
	"errors"

	"github.com/lskeey/go-filehub/config"
	"github.com/lskeey/go-filehub/internal/models"
	"github.com/lskeey/go-filehub/internal/repository"
	"github.com/lskeey/go-filehub/pkg/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repository.UserRepository
	cfg      config.Config
}

// NewAuthService creates a new authentication service.
func NewAuthService(repo *repository.UserRepository, cfg config.Config) *AuthService {
	return &AuthService{userRepo: repo, cfg: cfg}
}

// Register creates a new user account.
func (s *AuthService) Register(user *models.User) error {
	// Check if user already exists
	_, err := s.userRepo.FindUserByEmail(user.Email)
	if err == nil {
		return errors.New("user with this email already exists")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return errors.New("could not hash password")
	}
	user.Password = hashedPassword

	// Create user in the repository
	return s.userRepo.CreateUser(user)
}

// Login authenticates a user and returns a JWT.
func (s *AuthService) Login(email, password string) (string, error) {
	// Find user by email
	user, err := s.userRepo.FindUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid credentials")
		}
		return "", err
	}

	// Check password
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID, s.cfg.JWTSecretKey, s.cfg.JWTExpirationHours)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}
