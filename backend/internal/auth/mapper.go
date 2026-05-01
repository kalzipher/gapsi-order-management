package auth

func UserEntityToModel(entity UserEntity) User {
	return User{
		ID:           entity.ID,
		Email:        entity.Email,
		PasswordHash: entity.PasswordHash,
		Name:         entity.Name,
		CreatedAt:    entity.CreatedAt,
	}
}

func UserModelToEntity(user User) UserEntity {
	return UserEntity{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Name:         user.Name,
		CreatedAt:    user.CreatedAt,
	}
}

func RefreshTokenEntityToModel(entity RefreshTokenEntity) RefreshToken {
	return RefreshToken{
		ID:        entity.ID,
		UserID:    entity.UserID,
		TokenHash: entity.TokenHash,
		ExpiresAt: entity.ExpiresAt,
		RevokedAt: entity.RevokedAt,
		CreatedAt: entity.CreatedAt,
	}
}

func RefreshTokenModelToEntity(token RefreshToken) RefreshTokenEntity {
	return RefreshTokenEntity{
		ID:        token.ID,
		UserID:    token.UserID,
		TokenHash: token.TokenHash,
		ExpiresAt: token.ExpiresAt,
		RevokedAt: token.RevokedAt,
		CreatedAt: token.CreatedAt,
	}
}

func UserToDTO(user User) UserDTO {
	return UserDTO{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
}
