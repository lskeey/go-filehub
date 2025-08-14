package repository

import (
	"github.com/lskeey/go-filehub/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new user repository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser adds a new user to the database.
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// FindUserByEmail retrieves a user by their email address.
func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
