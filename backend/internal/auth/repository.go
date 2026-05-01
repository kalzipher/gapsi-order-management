package auth

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrRefreshTokenNotFound = errors.New("refresh token not found")
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	var entity UserEntity

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&entity).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	user := UserEntityToModel(entity)
	return &user, nil
}

func (r *Repository) FindUserByID(ctx context.Context, id string) (*User, error) {
	var entity UserEntity

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&entity).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	user := UserEntityToModel(entity)
	return &user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user User) error {
	entity := UserModelToEntity(user)

	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *Repository) SaveRefreshToken(ctx context.Context, token RefreshToken) error {
	entity := RefreshTokenModelToEntity(token)

	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *Repository) FindValidRefreshToken(ctx context.Context, tokenHash string) (*RefreshToken, error) {
	var entity RefreshTokenEntity

	err := r.db.WithContext(ctx).
		Where("token_hash = ?", tokenHash).
		Where("revoked_at IS NULL").
		Where("expires_at > ?", time.Now()).
		First(&entity).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRefreshTokenNotFound
		}

		return nil, err
	}

	token := RefreshTokenEntityToModel(entity)
	return &token, nil
}

func (r *Repository) RevokeRefreshToken(ctx context.Context, tokenHash string) error {
	return r.db.WithContext(ctx).
		Model(&RefreshTokenEntity{}).
		Where("token_hash = ?", tokenHash).
		Where("revoked_at IS NULL").
		Update("revoked_at", time.Now()).
		Error
}

func (r *Repository) RevokeAllUserRefreshTokens(ctx context.Context, userID string) error {
	return r.db.WithContext(ctx).
		Model(&RefreshTokenEntity{}).
		Where("user_id = ?", userID).
		Where("revoked_at IS NULL").
		Update("revoked_at", time.Now()).
		Error
}
