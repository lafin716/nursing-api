package persistence

import (
	"context"
	"github.com/google/uuid"
	"nursing_api/internal/domain/auth"
	"nursing_api/internal/domain/auth/repository"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	tokenEntity "nursing_api/pkg/ent/token"
	"time"
)

type authRepository struct {
	client *ent.TokenClient
	ctx    context.Context
}

func NewAuthRepository(dbClient *database.DatabaseClient) repository.AuthRepository {
	return &authRepository{
		client: dbClient.Client.Token,
		ctx:    dbClient.Ctx,
	}
}

func (u authRepository) SaveToken(userId uuid.UUID, token *auth.Token) (*auth.Token, error) {
	savedToken, err := u.client.Create().
		SetUserID(userId).
		SetAccessToken(token.AccessToken).
		SetAccessTokenExpires(time.Unix(token.AccessTokenExpires, 0)).
		SetRefreshToken(token.RefreshToken).
		SetRefreshTokenExpires(time.Unix(token.RefreshTokenExpires, 0)).
		Save(u.ctx)
	if err != nil {
		return nil, err
	}

	modelToken := toTokenModel(savedToken)
	return modelToken, nil
}

func (u authRepository) GetToken(userId uuid.UUID) (*auth.Token, error) {
	foundToken, err := u.client.Query().
		Where(
			tokenEntity.UserID(userId),
		).
		Only(u.ctx)
	if err != nil {
		return nil, err
	}

	return toTokenModel(foundToken), nil
}

func (u authRepository) DeleteToken(userId uuid.UUID) (bool, error) {
	_, err := u.client.Delete().
		Where(
			tokenEntity.UserID(userId),
		).
		Exec(u.ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func toTokenModel(entity *ent.Token) *auth.Token {
	return &auth.Token{
		AccessToken:         entity.AccessToken,
		AccessTokenExpires:  entity.AccessTokenExpires.Unix(),
		RefreshToken:        entity.RefreshToken,
		RefreshTokenExpires: entity.RefreshTokenExpires.Unix(),
	}
}

func toTokenEntity(model *auth.Token) *ent.Token {
	return &ent.Token{
		AccessToken:         model.AccessToken,
		AccessTokenExpires:  time.Unix(model.AccessTokenExpires, 0),
		RefreshToken:        model.RefreshToken,
		RefreshTokenExpires: time.Unix(model.RefreshTokenExpires, 0),
	}
}
