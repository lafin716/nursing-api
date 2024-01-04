package repository

import (
	"context"
	"nursing_api/ent"
	"nursing_api/ent/user"
)

type userRepository struct {
	client *ent.UserClient
	ctx    context.Context
}

type UserRepository interface {
	CreateUser(*ent.User) (*ent.User, error)
	CountUserByEmail(string) (int, error)
	GetUserByEmail(string) *ent.User
}

func NewUserRepository(client *ent.Client, ctx context.Context) UserRepository {
	return &userRepository{
		client: client.User,
		ctx:    ctx,
	}
}

func (u *userRepository) CreateUser(user *ent.User) (*ent.User, error) {
	savedUser, err := u.client.Create().
		SetUserName(user.UserName).
		SetUserEmail(user.UserEmail).
		SetUserPassword(user.UserPassword).
		SetCreatedAt(user.CreatedAt).
		Save(u.ctx)
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}

func (u *userRepository) CountUserByEmail(email string) (int, error) {
	client := u.client
	emailCount, err := client.
		Query().
		Where(
			user.UserEmail(email),
		).
		Count(u.ctx)
	if err != nil {
		return -1, err
	}

	return emailCount, nil
}

func (u *userRepository) GetUserByEmail(email string) *ent.User {
	client := u.client
	foundUser, err := client.Query().
		Where(
			user.UserEmail(email),
		).
		Only(u.ctx)
	if err != nil {
		return nil
	}

	return foundUser
}
