package auth

import (
	"time"

	"github.com/github/gapsi-order-management-dashboard/backend/internal/database"
)

type UserEntity struct {
	ID           string    `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	Email        string    `gorm:"column:email;type:text;not null;uniqueIndex"`
	PasswordHash string    `gorm:"column:password_hash;type:text;not null"`
	Name         string    `gorm:"column:name;type:text"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (UserEntity) TableName() string {
	return database.TableNameUsers
}

type RefreshTokenEntity struct {
	ID        string     `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    string     `gorm:"column:user_id;type:uuid;not null;index"`
	TokenHash string     `gorm:"column:token_hash;type:text;not null;index"`
	ExpiresAt time.Time  `gorm:"column:expires_at;not null"`
	RevokedAt *time.Time `gorm:"column:revoked_at"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime"`

	User UserEntity `gorm:"foreignKey:UserID"`
}

func (RefreshTokenEntity) TableName() string {
	return database.TableNameRefreshTokens
}
