package service

import (
	"context"
	"os"
	"time"

	"go-pos/internal/core/port"
	"go-pos/internal/core/util"
)

/**
 * AuthService implements port.AuthService interface
 * and provides an access to the user repository
 * and token service
 */
type AuthService struct {
	repo port.UserRepository
	ts   port.TokenService
}

// NewAuthService creates a new auth service instance
func NewAuthService(repo port.UserRepository, ts port.TokenService) *AuthService {
	return &AuthService{
		repo,
		ts,
	}
}

// Login gives a registered user an access token if the credentials are valid
func (as *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := as.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", port.ErrInvalidCredentials
	}

	err = util.ComparePassword(password, user.Password)
	if err != nil {
		return "", port.ErrInvalidCredentials
	}

	accessTokenDurationStr := os.Getenv("TOKEN_DURATION")
	accessTokenDuration, err := time.ParseDuration(accessTokenDurationStr)
	if err != nil {
		return "", err
	}

	accessToken, err := as.ts.CreateToken(user, accessTokenDuration)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
