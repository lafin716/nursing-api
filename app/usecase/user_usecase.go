package usecase

import (
	"nursing_api/app/models"
	"nursing_api/app/repository"
	"nursing_api/ent"
)

type UserGateway struct {
	userRepository repository.UserRepository
}

func NewUserGateway(userRepository repository.UserRepository) *UserGateway {
	return &UserGateway{
		userRepository: userRepository,
	}
}

type UserService struct {
	gateway *UserGateway
}

type UserUseCase interface {
	SaveUser(user *models.User) *models.User
	GetUserByEmail(email string) *models.User
	CountUserByEmail(email string) int
}

func NewUserService(gateway *UserGateway) UserUseCase {
	return &UserService{
		gateway: gateway,
	}
}

func (s *UserService) SaveUser(user *models.User) *models.User {
	entityUser := toEntity(user)
	savedUser, err := s.gateway.userRepository.CreateUser(entityUser)
	if err != nil {
		return nil
	}

	return toModel(savedUser)
}

func (s *UserService) GetUserByEmail(email string) *models.User {
	entityUser := s.gateway.userRepository.GetUserByEmail(email)
	modelUser := toModel(entityUser)
	return modelUser
}

func (s *UserService) CountUserByEmail(email string) int {
	count, err := s.gateway.userRepository.CountUserByEmail(email)
	if err != nil {
		return -1
	}

	return count
}

func toModel(entity *ent.User) *models.User {
	return &models.User{
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

func toEntity(model *models.User) *ent.User {
	return &ent.User{
		ID:           model.ID,
		UserName:     model.Name,
		UserEmail:    model.Email,
		UserPassword: model.PasswordHash,
		UserStatus:   model.UserStatus,
		UserType:     model.UserType,
		FailCount:    model.FailCount,
		LastSignedIn: model.LastLoggedIn,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}
}
