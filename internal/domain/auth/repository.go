package auth

import (
	"context"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	tokenEntity "nursing_api/pkg/ent/token"
	"time"
)

type authRepository struct {
	client *ent.TokenClient
	ctx    context.Context
}

func NewRepository(dbClient *database.DatabaseClient) AuthRepository {
	return &authRepository{
		client: dbClient.Client.Token,
		ctx:    dbClient.Ctx,
	}
}

func (u authRepository) SaveToken(userId uuid.UUID, token *Token) (*Token, error) {
	savedToken, err := u.client.Create().
		SetUserID(userId).
		SetAccessToken(token.AccessToken).
		SetAccessTokenExpires(time.Unix(token.AccessTokenExpires, 0)).
		SetRefreshToken(token.RefreshToken).
		SetRefreshTokenExpires(time.Unix(token.RefreshTokenExpires, 0)).
		SetAutoLogin(token.AutoLogin).
		Save(u.ctx)
	if err != nil {
		return nil, err
	}

	modelToken := toTokenModel(savedToken)
	return modelToken, nil
}

func (u authRepository) GetToken(userId uuid.UUID) (*Token, error) {
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

func toTokenModel(entity *ent.Token) *Token {
	return &Token{
		AccessToken:         entity.AccessToken,
		AccessTokenExpires:  entity.AccessTokenExpires.Unix(),
		RefreshToken:        entity.RefreshToken,
		RefreshTokenExpires: entity.RefreshTokenExpires.Unix(),
		AutoLogin:           entity.AutoLogin,
	}
}

func toTokenEntity(model *Token) *ent.Token {
	return &ent.Token{
		AccessToken:         model.AccessToken,
		AccessTokenExpires:  time.Unix(model.AccessTokenExpires, 0),
		RefreshToken:        model.RefreshToken,
		RefreshTokenExpires: time.Unix(model.RefreshTokenExpires, 0),
		AutoLogin:           model.AutoLogin,
	}
}
