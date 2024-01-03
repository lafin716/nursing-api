package repository

import (
	"context"
	"nursing_api/ent"
	"nursing_api/ent/user"
)

type userRepository struct {
	client *ent.UserClient
}

type UserRepository interface {
	CountUserByEmail(email string) (int, error)
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{client: client.User}
}

func (u *userRepository) CountUserByEmail(email string) (int, error) {
	client, err := GetClient()
	if err != nil {
		return -1, err
	}
	defer client.Close()
	ctx := context.Background()

	emailCount, err := client.User.
		Query().
		Where(
			user.UserEmail(email),
		).
		Count(ctx)

	return emailCount, nil
}
