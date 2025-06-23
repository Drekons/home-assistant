package auth

import (
	"context"
	"fmt"
	"github.com/Drekons/home-assistant/backend/cmd/app"
	"github.com/Drekons/home-assistant/backend/internal/entity"
	"github.com/Drekons/home-assistant/backend/internal/interfaces"
	"regexp"
	"strings"
)

type Registry struct {
	repo interfaces.Repository
}

func NewRegistry(deps *app.Deps) *Registry {
	return &Registry{repo: deps.Repo()}
}

func (r *Registry) Register(ctx context.Context, username, password, email string) (*entity.User, error) {
	// Check if username already exists
	user, err := r.repo.Users().FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("error finding user by email: %w", err)
	}
	if user != nil {
		return nil, fmt.Errorf("email already exists")
	}

	// Check if email is in a valid format
	if !r.isValidEmail(email) {
		return nil, fmt.Errorf("invalid email format")
	}

	// Check if username is alphanumeric and between 8 and 20 characters long
	if !isUsernameValid(username) {
		return nil, fmt.Errorf("invalid username format")
	}

	// Check if password is between 8 and 20 characters long and contains at least one uppercase letter, one lowercase letter, one number, and one special character
	if !isPasswordValid(password) {
		return nil, fmt.Errorf("invalid password format")
	}

	// Create new user
	newUser := &entity.User{
		Username: username,
		Email:    email,
	}

	err = newUser.SetPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error setting password")
	}

	err = r.repo.Users().Create(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return newUser, nil
}

func isUsernameValid(username string) bool {
	// Check if username length is between 3 and 20 characters
	if len(username) < 3 || len(username) > 20 {
		return false
	}

	// Check if username contains only alphanumeric characters
	for _, char := range username {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			return false
		}
	}

	return true
}

func isPasswordValid(password string) bool {
	// Check if password is between 8 and 20 characters long
	if len(password) < 8 || len(password) > 20 {
		return false
	}

	// Check if password contains at least one uppercase letter, one lowercase letter, one number, and one special character
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		case strings.ContainsRune("!@#$%^&*()-_=+[]{}|;:'\",.<>?/", char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}

func (r *Registry) isValidEmail(s string) bool {
	// Regular expression pattern for email validation
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	// Check if the email matches the pattern
	return emailRegex.MatchString(strings.ToLower(s))
}
