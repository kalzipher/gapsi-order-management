package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

type Service struct {
	repo RepositoryPort
	jwt  TokenPort
}

func NewService(repo RepositoryPort, jwt TokenPort) *Service {
	return &Service{
		repo: repo,
		jwt:  jwt,
	}
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	email := normalizeEmail(req.Email)

	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}

		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	accessToken, _, err := s.jwt.GenerateAccessToken(*user)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshExpiresAt, err := s.jwt.GenerateRefreshToken(*user)
	if err != nil {
		return nil, err
	}

	refreshTokenHash := HashToken(refreshToken)

	err = s.repo.SaveRefreshToken(ctx, RefreshToken{
		UserID:    user.ID,
		TokenHash: refreshTokenHash,
		ExpiresAt: refreshExpiresAt,
	})
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         UserToDTO(*user),
	}, nil
}

func (s *Service) Refresh(ctx context.Context, req RefreshRequest) (*RefreshResponse, error) {
	claims, err := s.jwt.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	tokenHash := HashToken(req.RefreshToken)

	storedToken, err := s.repo.FindValidRefreshToken(ctx, tokenHash)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	if storedToken.UserID != claims.UserID {
		return nil, ErrInvalidRefreshToken
	}

	user, err := s.repo.FindUserByID(ctx, claims.UserID)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	accessToken, _, err := s.jwt.GenerateAccessToken(*user)
	if err != nil {
		return nil, err
	}

	return &RefreshResponse{
		AccessToken: accessToken,
	}, nil
}

func (s *Service) Logout(ctx context.Context, req LogoutRequest) error {
	tokenHash := HashToken(req.RefreshToken)
	return s.repo.RevokeRefreshToken(ctx, tokenHash)
}

func (s *Service) EnsureAdminUser(ctx context.Context, email, password, name string) error {
	email = normalizeEmail(email)

	_, err := s.repo.FindUserByEmail(ctx, email)
	if err == nil {
		return nil
	}

	if !errors.Is(err, ErrUserNotFound) {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := User{
		Email:        email,
		PasswordHash: string(passwordHash),
		Name:         name,
	}

	return s.repo.CreateUser(ctx, user)
}

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
