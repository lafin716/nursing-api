package user

import (
	"context"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	tokenEntity "nursing_api/pkg/ent/token"
	userEntity "nursing_api/pkg/ent/user"
)

type userRepository struct {
	root   *ent.Client
	client *ent.UserClient
	ctx    context.Context
}

type Repository interface {
	CreateUser(user *User) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(userId uuid.UUID) (*User, error)
	CountUserByEmail(email string) (int, error)
	Delete(userId uuid.UUID) (bool, error)
}

func NewRepository(dbClient database.DatabaseClient) Repository {
	return &userRepository{
		root:   dbClient.GetClient(),
		client: dbClient.GetClient().User,
		ctx:    dbClient.GetCtx(),
	}
}

func (u userRepository) GetUserById(userId uuid.UUID) (*User, error) {
	foundUser, err := u.client.Query().Where(userEntity.ID(userId)).Only(u.ctx)
	if err != nil {
		return nil, err
	}

	return toUserDomain(foundUser), nil
}

func (u userRepository) CreateUser(user *User) (*User, error) {
	entity := toUserEntity(user)
	savedUser, err := u.client.Create().
		SetUserName(entity.UserName).
		SetUserEmail(entity.UserEmail).
		SetUserPassword(entity.UserPassword).
		SetCreatedAt(entity.CreatedAt).
		Save(u.ctx)
	if err != nil {
		return nil, err
	}

	return toUserDomain(savedUser), nil
}

func (u userRepository) GetUserByEmail(email string) (*User, error) {
	client := u.client
	foundUser, err := client.Query().
		Where(
			userEntity.UserEmail(email),
		).
		Only(u.ctx)
	if err != nil {
		return nil, err
	}

	return toUserDomain(foundUser), nil
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

func (u userRepository) Delete(userId uuid.UUID) (bool, error) {
	// 의존 관계에 있는 토큰 테이블도 유저토큰 삭제처리를 한다. (예외처리 없음)
	u.root.Token.Delete().Where(tokenEntity.UserID(userId)).Exec(u.ctx)
	err := u.client.DeleteOneID(userId).Exec(u.ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func toUserDomain(entity *ent.User) *User {
	return &User{
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

func toUserEntity(domain *User) *ent.User {
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
