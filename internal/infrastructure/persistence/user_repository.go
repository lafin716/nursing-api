package persistence

import (
	"context"
	"nursing_api/internal/domain/user"
	"nursing_api/internal/domain/user/repository"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	userEntity "nursing_api/pkg/ent/user"
)

type userRepository struct {
	client *ent.UserClient
	ctx    context.Context
}

func NewUserRepository(dbClient *database.DatabaseClient) repository.UserRepository {
	return &userRepository{
		client: dbClient.Client.User,
		ctx:    dbClient.Ctx,
	}
}

func (u userRepository) CreateUser(user *user.User) (*user.User, error) {
	entity := toEntity(user)
	savedUser, err := u.client.Create().
		SetUserName(entity.UserName).
		SetUserEmail(entity.UserEmail).
		SetUserPassword(entity.UserPassword).
		SetCreatedAt(entity.CreatedAt).
		Save(u.ctx)
	if err != nil {
		return nil, err
	}

	return toDomain(savedUser), nil
}

func (u userRepository) GetUserByEmail(email string) (*user.User, error) {
	client := u.client
	foundUser, err := client.Query().
		Where(
			userEntity.UserEmail(email),
		).
		Only(u.ctx)
	if err != nil {
		return nil, err
	}

	return toDomain(foundUser), nil
}

func (u userRepository) CountUserByEmail(email string) (int, error) {
	client := u.client
	emailCount, err := client.
		Query().
		Where(
			userEntity.UserEmail(email),
		).
		Count(u.ctx)
	if err != nil {
		return -1, err
	}

	return emailCount, nil
}

func toDomain(entity *ent.User) *user.User {
	return &user.User{
		ID:           entity.ID,
		Name:         entity.UserName,
		Email:        entity.UserEmail,
		PasswordHash: entity.UserPassword,
		UserStatus:   entity.UserStatus,
		UserType:     entity.UserType,
		FailCount:    entity.FailCount,
		LastLoggedIn: entity.LastSignedIn,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}

func toEntity(domain *user.User) *ent.User {
	return &ent.User{
		ID:           domain.ID,
		UserName:     domain.Name,
		UserEmail:    domain.Email,
		UserPassword: domain.PasswordHash,
		UserStatus:   domain.UserStatus,
		UserType:     domain.UserType,
		FailCount:    domain.FailCount,
		LastSignedIn: domain.LastLoggedIn,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
